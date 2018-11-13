package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64 //index of blocks
	Timestamp    int64 //the time when the block create
	SelfHash     string
	PreBlockHash string //the hash of last block
	Data         string //the data of the block
}

func calcBlockHash(block Block) string {
	toBeHashed := string(block.Index) + string(block.Timestamp) + block.PreBlockHash + block.Data
	hashInBytes := sha256.Sum256([]byte(toBeHashed))
	return hex.EncodeToString(hashInBytes[:])
}

func GenerateNewBlock(lastBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Data = data
	newBlock.Index = lastBlock.Index + 1
	newBlock.PreBlockHash = lastBlock.SelfHash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.SelfHash = calcBlockHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	lastBlock := Block{}
	lastBlock.Index = -1
	lastBlock.SelfHash = ""
	return GenerateNewBlock(lastBlock, "Genesis Block")

}
