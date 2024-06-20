package main

import (
	"coin/cli"
	"coin/db"
)

func main() {
	defer db.Close()
	// blockchain.Blockchain()
	// blockchain.Blockchain().AddBlock("test")
	// block, _ := blockchain.FindBlock("556936855dc733e36c85ee69d4b6aaac7dab981ebb5c4bbfdb6db41cffc73c00")
	// fmt.Println("Data", block.Data)
	// fmt.Println("Hash", block.Hash)
	// fmt.Println("PrevHash", block.PrevHash)
	// fmt.Println("Height", block.Height)
	cli.Start()
}
