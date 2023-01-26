package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func (b *Block) Print() {
	fmt.Printf("Timestamp: %v\n", b.timestamp)
	fmt.Printf("Nonce: %v\n", b.nonce)
	fmt.Printf("Previous Hash: %s\n", b.previousHash)
	fmt.Printf("Transactions: %v\n", b.transactions)

}

func NewBlock(nonce int, previousHash string) *Block {
	var block Block
	block.timestamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = previousHash
	return &block
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	var chain Blockchain
	chain.CreateBlock(0, "Init Hash")
	return &chain
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
}

func init() {
	log.SetPrefix("Blockchain:")
}

func main() {
	blockchain := NewBlockchain()
	blockchain.CreateBlock(5, "hash 1")
	blockchain.CreateBlock(2, "hash 2")
	blockchain.Print()
}
