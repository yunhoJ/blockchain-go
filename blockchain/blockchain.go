package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type blockChain struct {
	blocks []*Block
}

// 싱글톤 패턴 하나의 인스턴스만을 공유하는 방법
var b *blockChain

var Once sync.Once

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockchain().blocks) + 1}
	newBlock.calculateHash()
	return &newBlock

}
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash

}
func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	hexHash := fmt.Sprintf("%x", hash)
	b.Hash = hexHash
}
func (b *blockChain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}
func GetBlockchain() *blockChain {
	//b초기화 확인
	if b == nil {
		Once.Do(func() { //한번만 실행하도록 함
			b = &blockChain{} // 구조체의 주소값을 할당함
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockChain) AllBlock() []*Block {
	return b.blocks
}

var ErrNotFound = errors.New("block not found")

func (b *blockChain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil

}
