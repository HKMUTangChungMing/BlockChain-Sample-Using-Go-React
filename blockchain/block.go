package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// 定義區塊結構
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// 計算區塊哈希值
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

// 生成新區塊
func GenerateBlock(prevBlock Block, data string) Block {
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}
