package main

import (
	"fmt"
	"github.com/johnmcdnl/johncoin"
	"strconv"
)

func main() {
	bc := johncoin.NewBlockchain()

	bc.AddBlock("Send 1 JC to Ivan")
	bc.AddBlock("Send 2 more JC to Ivan")
	bc.AddBlock("Send 1 JC to John")
	bc.AddBlock("Send 10 JC to Jim")
	bc.AddBlock("Send 4 more JC to Jim")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := johncoin.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
	fmt.Println()
}

