package chain

import (
	"GoBlockchain/Transaction"
	"GoBlockchain/block"
	"GoBlockchain/utils"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Blockchain struct {
	transactionPool   []*Transaction.Transaction
	chain             []*block.Block
	blockchainAddress string
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *block.Block {
	b := block.NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction.Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, blk := range bc.chain {
		fmt.Printf("%s block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		blk.Print()
	}
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32, senderPublicKey *ecdsa.PublicKey, s *utils.Signature) bool {
	t := Transaction.NewTransaction(sender, recipient, value)

	if sender == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	if bc.VerifyTransactionSignature(senderPublicKey, s, t) {
		//if bc.CalculateTotalAmount(sender) < value {
		//	log.Println("ERROR: Not Enough Balance in a wallet")
		//	return false
		//}
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	log.Println("ERROR: Verify Transaction")
	return false
}

func (bc *Blockchain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, sign *utils.Signature, t *Transaction.Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, h[:], sign.R, sign.S)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction.Transaction {
	transactions := make([]*Transaction.Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, Transaction.NewTransaction(t.GetTransactionDetails()))
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction.Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := block.DummyBlock(0, nonce, previousHash, transactions)
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	var nonce int = 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=Mining, Status=Success")
	return true
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0
	for _, b := range bc.chain {
		for _, t := range b.GetTransactions() {
			senderAddress, receiverAddress, value := t.GetTransactionDetails()
			if receiverAddress == blockchainAddress {
				totalAmount += value
			}
			if blockchainAddress == senderAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	var blk *block.Block = &block.Block{}
	var chain Blockchain
	chain.blockchainAddress = blockchainAddress
	chain.CreateBlock(0, blk.Hash())
	return &chain
}
