package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)



//to sta rt the main program
func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Println("")
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Println("")
		fmt.Printf("Hash: %x\n", block.Data)
	}
}
