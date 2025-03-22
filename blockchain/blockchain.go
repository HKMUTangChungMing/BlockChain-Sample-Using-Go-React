package blockchain

import (
	"time"
)

var Blockchain []Block

// 創建創世區塊
func CreateGenesisBlock() Block {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}

// 初始化區塊鏈
func InitBlockchain() {
	Blockchain = append(Blockchain, CreateGenesisBlock())
}
