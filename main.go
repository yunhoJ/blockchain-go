package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockChain struct {
	blocks []block
}

func (b *blockChain) getprevhash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}
func (b *blockChain) addBlock(data string) {

	newblock := block{data, "", b.getprevhash()}
	hash := sha256.Sum256([]byte(newblock.data + newblock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	newblock.hash = hexHash
	b.blocks = append(b.blocks, newblock)
}
func (b *blockChain) listBlocks() {
	for _, data := range b.blocks {
		fmt.Println(data)
	}
}
func main() {
	chain := blockChain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}
