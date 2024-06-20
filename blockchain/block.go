package blockchain

import (
	"coin/db"
	"coin/utils"
	"crypto/sha256"
	"errors"
	"fmt"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

//	func (b *Block) toBytes() []byte {
//		var blockBuffer bytes.Buffer
//		encoder := gob.NewEncoder(&blockBuffer)
//		err := encoder.Encode(b)
//		utils.HandleErr(err)
//		return blockBuffer.Bytes()
//	}
func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}
func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		data,
		"",
		prevHash,
		height,
	}
	payload := block.Data + block.Hash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist() // db에 저장
	return block
}

var ErrNotFound = errors.New("block not found")

func FindBlock(hash string) (*Block, error) {

	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}

	b := &Block{"", "", "", 0}
	b.restore(blockBytes)
	return b, nil
}
func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}
