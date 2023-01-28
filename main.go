package main

import (
	chain "GoBlockchain/chain"
	"log"
)

func init() {
	log.SetPrefix("Blockchain:")
}

func main() {
	myBlockchainAddress := "my_blockchain_address"
	blockchain := chain.NewBlockchain(myBlockchainAddress)

	blockchain.AddTransaction("A", "B", 1.0)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("C", "B", 2.0)
	blockchain.AddTransaction("A", "C", 3.0)
	blockchain.Mining()
	blockchain.Print()
}
