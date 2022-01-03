package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int    // 데이터 레코드의 위치
	Timestamp string // 데이터가 작성될때의 시간이 작성되며 자동으로 결정
	BPM       int    // Beats For Minute, 이것은 pulse rate을 의미
	Hash      string // SHA-256을 이용하며 이 데이터 레코드의 식별을 하는데 사용
	PrevHash  string // 이전 데이터 레코드의 Hash를 의미
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
