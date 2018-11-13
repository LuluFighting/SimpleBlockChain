package core

import (
	"fmt"
	"log"
)

type BlockChain struct {
	BlockList []*Block
}

func CreatBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	bc := BlockChain{}
	bc.AddNewBlock(&genesisBlock)
	return &bc
}

func (bc *BlockChain) Print() {
	for _, block := range bc.BlockList {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("SelfHash: %s\n", block.SelfHash)
		fmt.Printf("PreBlockHash: %s\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("\n")
	}
}
func (bc *BlockChain) SendData(data string) {
	lastBlock := bc.BlockList[len(bc.BlockList)-1]
	newBlock := GenerateNewBlock(*lastBlock, data)
	bc.AddNewBlock(&newBlock)

}
func (bc *BlockChain) AddNewBlock(block *Block) {
	if len(bc.BlockList) == 0 {
		bc.BlockList = append(bc.BlockList, block)
		return
	}
	if IsValid(*block, *bc.BlockList[len(bc.BlockList)-1]) {
		bc.BlockList = append(bc.BlockList, block)
	} else {
		log.Fatal("invalid block")
	}
}

func IsValid(newBlock, preBlock Block) bool {
	if newBlock.Index != preBlock.Index+1 {
		return false
	}
	if newBlock.PreBlockHash != preBlock.SelfHash {
		return false
	}
	if newBlock.SelfHash != calcBlockHash(newBlock) {
		return false
	}
	return true
}
