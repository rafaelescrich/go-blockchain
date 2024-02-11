package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	Nonce        int
	PrevHash     string
	Hash         string
	MerkleRoot   string
}

type Blockchain struct {
	blocks []*Block
}

var blockchain = Blockchain{}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.MerkleRoot + block.PrevHash + strconv.Itoa(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func calculateMerkleRoot(transactions []Transaction) string {
	var merkleTree []string

	for _, t := range transactions {
		merkleTree = append(merkleTree, t.Hash())
	}

	for i := 0; len(merkleTree) > 1; {
		if i+1 == len(merkleTree) {
			merkleTree = append(merkleTree, merkleTree[i]) // duplicate last hash if odd number of hashes
		}
		merged := merkleTree[i] + merkleTree[i+1]
		h := sha256.New()
		h.Write([]byte(merged))
		merkleTree = append(merkleTree[2:], hex.EncodeToString(h.Sum(nil)))
		i += 2
	}

	if len(merkleTree) == 0 {
		return ""
	}
	return merkleTree[0]
}

func (t *Transaction) Hash() string {
	record := t.Sender + t.Recipient + fmt.Sprintf("%f", t.Amount)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (bc *Blockchain) addBlock(block *Block) {
	if len(bc.blocks) > 0 {
		block.PrevHash = bc.blocks[len(bc.blocks)-1].Hash
	}
	block.Hash = calculateHash(*block)
	bc.blocks = append(bc.blocks, block)
}

func mineBlock(block *Block, difficulty int) {
	block.MerkleRoot = calculateMerkleRoot(block.Transactions)
	for !isHashValid(block.Hash, difficulty) {
		block.Nonce++
		block.Hash = calculateHash(*block)
	}
	fmt.Println("Block mined:", block.Hash)
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func main() {
	portPtr := flag.String("p", "8080", "Port on which the server will run")
	flag.Parse()

	r := gin.Default()

	r.POST("/transactions/new", func(c *gin.Context) {
		var transaction Transaction
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(400, gin.H{"message": "Bad request"})
			return
		}

		// Assuming we're adding to the latest block (simplified)
		if len(blockchain.blocks) == 0 {
			blockchain.blocks = append(blockchain.blocks, &Block{})
		}
		blockchain.blocks[len(blockchain.blocks)-1].Transactions = append(blockchain.blocks[len(blockchain.blocks)-1].Transactions, transaction)
		c.JSON(201, transaction)
	})

	r.GET("/mine", func(c *gin.Context) {
		lastBlock := blockchain.blocks[len(blockchain.blocks)-1]
		newBlock := &Block{
			Index:        len(blockchain.blocks),
			Timestamp:    time.Now().String(),
			Transactions: lastBlock.Transactions,
		}

		mineBlock(newBlock, 4) // Difficulty level is 4
		blockchain.addBlock(newBlock)
		c.JSON(200, newBlock)
	})

	r.Run(":" + *portPtr) // listen and serve on the specified port

}
