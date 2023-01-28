package main

import (
	"GoBlockchain/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain:")
}

func main() {
	//myBlockchainAddress := "my_blockchain_address"
	//blockchain := chain.NewBlockchain(myBlockchainAddress)
	//
	//blockchain.AddTransaction("A", "B", 1.0)
	//blockchain.Mining()
	//blockchain.Print()
	//
	//blockchain.AddTransaction("C", "B", 2.0)
	//blockchain.AddTransaction("A", "C", 3.0)
	//blockchain.Mining()
	//blockchain.Print()
	//
	//fmt.Printf("My %.1f\n", blockchain.CalculateTotalAmount("my_blockchain_address"))
	//fmt.Printf("C %.1f\n", blockchain.CalculateTotalAmount("C"))

	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())
}
