package blockchain

import (
	"coin/db"
	"coin/utils"
	"errors"
	"strings"
	"time"
)

// const difficulty int = 2

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty`
	Nonce      int    `json:"nonce`
	Timestamp  int    `json:"timestamp`
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
func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	b.Timestamp = int(time.Now().Unix())
	for {
		hash := utils.Hash(b)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}

	}
}
func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		data,
		"",
		prevHash,
		height,
		Blockchain().difficulty(),
		0,
		0,
	}
	block.mine()
	block.persist() // db에 저장
	return block
}

var ErrNotFound = errors.New("block not found")

func FindBlock(hash string) (*Block, error) {

	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}

	b := &Block{"", "", "", 0, 0, 0, 0}
	b.restore(blockBytes)
	return b, nil
}
func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}
