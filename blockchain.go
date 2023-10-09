package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         CalculateHash(fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)),
	}
	return block
}

func DisplayBlocks(bc *Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s, Nonce: %d, Previous Hash: %s, Current Hash: %s\n", block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
	}
}

func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.Hash = CalculateHash(fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash))
}

func VerifyChain(bc *Blockchain) bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != CalculateHash(fmt.Sprintf("%s%d%s", currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash)) {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func CalculateHash(stringToHash string) string {
	hashInBytes := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hashInBytes[:])
}
