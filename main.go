package main

import (
	"coin/blockchain"
	"coin/cli"
	"coin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	// // blockchain.Blockchain().AddBlock("test")
	// // block, _ := blockchain.FindBlock("556936855dc733e36c85ee69d4b6aaac7dab981ebb5c4bbfdb6db41cffc73c00")
	// // fmt.Println("Data", block.Data)
	// // fmt.Println("Hash", block.Hash)
	// // fmt.Println("PrevHash", block.PrevHash)
	// // fmt.Println("Height", block.Height)
	cli.Start()
	// difficulty := 3
	// target := strings.Repeat("0", difficulty)
	// fmt.Println(target)
	// nonce := 1
	// for {
	// 	hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
	// 	if strings.HasPrefix(hash, target) {
	// 		fmt.Println(hash, nonce)
	// 		return
	// 	} else {
	// 		nonce++
	// 	}
	// }
}
