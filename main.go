package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("1 BTC")
	bc.AddBlock("2 BTC")

	for _, block := range bc.blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}

// TODO: add
// 1. Improved hashing algorithm
// 2. Proof-of-Work
// 3. Consensus
// 4. Transactions
