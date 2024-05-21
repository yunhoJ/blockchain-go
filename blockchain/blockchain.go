package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockChain struct {
	blocks []*Block
}

// 싱글톤 패턴 하나의 인스턴스만을 공유하는 방법
var b *blockChain

var Once sync.Once

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
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
