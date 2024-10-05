package main

import (
	assignment01bca "blockchain/assignment01bca"
	"fmt"
)

func main() {

	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "")
	assignment01bca.Blockchain = append(assignment01bca.Blockchain, *genesisBlock)

	block1 := assignment01bca.NewBlock("Alice pays Bob", 1, genesisBlock.Hash)
	assignment01bca.Blockchain = append(assignment01bca.Blockchain, *block1)

	block2 := assignment01bca.NewBlock("Bob pays Charlie", 2, block1.Hash)
	assignment01bca.Blockchain = append(assignment01bca.Blockchain, *block2)

	block3 := assignment01bca.NewBlock("Charlie pays Dave", 3, block2.Hash)
	assignment01bca.Blockchain = append(assignment01bca.Blockchain, *block3)

	fmt.Println("Blockchain before modification:")
	assignment01bca.ListBlocks()

	fmt.Println("Modifying block 1's transaction...")
	assignment01bca.ChangeBlock(1, "Alice pays Eve")

	fmt.Println("Blockchain after modification:")
	assignment01bca.ListBlocks()

	fmt.Println("Verifying the blockchain's integrity:")
	if assignment01bca.VerifyChain() {
		fmt.Println("The blockchain is valid.")
	} else {
		fmt.Println("The blockchain has been tampered with!")
	}
}
