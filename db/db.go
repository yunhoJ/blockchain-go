package db

import (
	"coin/utils"
	"sync"

	"github.com/boltdb/bolt"
)

// bolt = key, value db
// bucket = table , data를 구분 해서 저장 할수 있음
// 싱글톤 패턴
var db *bolt.DB

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
	checkpoint   = "checkpoint"
)

var Once sync.Once

func DB() *bolt.DB {
	Once.Do(func() {
		//init db
		dbPointer, err := bolt.Open(dbName, 0600, nil) // 디비 생성
		utils.HandleErr(err)
		db = dbPointer
	})
	if db == nil {
		var err error
		// read-write transaction
		err = db.Update(func(tx *bolt.Tx) error {

			_, err = tx.CreateBucket([]byte(dataBucket))
			_, err = tx.CreateBucket([]byte(blocksBucket))

			return err
		})
		utils.HandleErr(err)
	}
	return db
}

// key ,value 저장 - byte 만 저장 가능 함
func SaveBlock(hash string, data []byte) {
	// fmt.Printf("save hash : %s \n data %b", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)

}

func GetCheckpoint() []byte {
	var data []byte
	err := DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	utils.HandleErr(err)
	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
func Close() {
	DB().Close()

}
