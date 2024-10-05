package assignment01bca

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

var Blockchain []Block

func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	blockData := transaction + string(nonce) + previousHash
	blockHash := CalculateHash(blockData)

	newBlock := Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         blockHash,
	}

	return &newBlock
}

func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("\tTransaction: %s\n", block.Transaction)
		fmt.Printf("\tNonce: %d\n", block.Nonce)
		fmt.Printf("\tPrevious Hash: %s\n", block.PreviousHash)
		fmt.Printf("\tCurrent Hash: %s\n", block.Hash)
		fmt.Println()
	}
}

func ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex < 0 || blockIndex >= len(Blockchain) {
		fmt.Println("Invalid block index")
		return
	}

	Blockchain[blockIndex].Transaction = newTransaction
	Blockchain[blockIndex].Hash = CalculateHash(newTransaction + string(Blockchain[blockIndex].Nonce) + Blockchain[blockIndex].PreviousHash)
}

func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		previousBlock := Blockchain[i-1]
		currentBlock := Blockchain[i]

		calculatedHash := CalculateHash(currentBlock.Transaction + string(currentBlock.Nonce) + currentBlock.PreviousHash)
		if currentBlock.Hash != calculatedHash || currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Println("Blockchain is invalid!")
			return false
		}
	}
	fmt.Println("Blockchain is valid.")
	return true
}
