package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// 어떤 타입이든 받을수 있음
func ToBytes(i interface{}) []byte {
	var Buffer bytes.Buffer
	encoder := gob.NewEncoder(&Buffer)
	err := encoder.Encode(i)
	HandleErr(err)
	return Buffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(i)
	HandleErr(err)
}

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i) //%v는 기본 formmater ,
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}
