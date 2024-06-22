package blockchain

import (
	"coin/db"
	"coin/utils"
	"fmt"
	"sync"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5 // 블록 5개마다 난이도 검사
	blockInterver      int = 2 // 2분에 1개씩 생성
	allowdRange        int = 2
)

type blockChain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:currentDifficulty`
}

// 싱글톤 패턴 하나의 인스턴스만을 공유하는 방법
var b *blockChain
var Once sync.Once

func Blockchain() *blockChain {
	//b초기화 확인
	if b == nil {
		Once.Do(func() { //한번만 실행하도록 함
			b = &blockChain{Height: 0,
				CurrentDifficulty: 0} // 구조체의 주소값을 할당함
			fmt.Printf("hash %s, height %d\n", b.NewestHash, b.Height)
			persitedBlockchin := db.GetCheckpoint()
			if persitedBlockchin == nil {
				// db 내용이 없으면 생성
				b.AddBlock("Genesis")
			} else {
				fmt.Println("restore....")
				b.restore(persitedBlockchin)
				// 마지막 블록 복원
			}
			fmt.Printf("hash %s, height %d\n", b.NewestHash, b.Height)

		})
	}
	return b
}

func (b *blockChain) difficulty() int {
	// b.recalculateDifficulty()
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}
func (b *blockChain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]

	actualTime := (newestBlock.Timestamp - lastRecalculatedBlock.Timestamp) / 60
	expectTime := difficultyInterval * blockInterver

	//8~12분 허용
	if actualTime < (expectTime - allowdRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectTime + allowdRange) {
		return b.CurrentDifficulty - 1
	} else {
		return b.CurrentDifficulty
	}
	// return b.CurrentDifficulty + 1
}

func (b *blockChain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)

	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
	// block := Block{data, "", b.NewestHash, }
}

func (b *blockChain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}
func (b *blockChain) restore(data []byte) {
	utils.FromBytes(b, data)
}

// 전체 블록 찾는다
func (b *blockChain) Blocks() []*Block {
	var blockList []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blockList = append(blockList, block)
		hashCursor = block.PrevHash
		if hashCursor == "" {
			break
		}
	}
	return blockList
}
