package block

import (
	"GoBlockchain/Transaction"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction.Transaction
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64                      `json:"Timestamp"`
		Nonce        int                        `json:"Nonce"`
		PreviousHash [32]byte                   `json:"PreviousHash"`
		Transactions []*Transaction.Transaction `json:"Transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("Timestamp: %v\n", b.timestamp)
	fmt.Printf("Nonce: %v\n", b.nonce)
	fmt.Printf("Previous Hash: %x\n", b.previousHash)

	for _, transaction := range b.transactions {
		transaction.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) GetTransactions() []*Transaction.Transaction {
	return b.transactions
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction.Transaction) *Block {
	var block Block
	block.timestamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = previousHash
	block.transactions = transactions
	return &block
}

func DummyBlock(timeStamp int64, nonce int, previousHash [32]byte, transactions []*Transaction.Transaction) *Block {
	var block Block
	block.timestamp = timeStamp
	block.nonce = nonce
	block.previousHash = previousHash
	block.transactions = transactions
	return &block
}
