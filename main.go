package main

// Go looks nice. 🐘

import (
	"bytes"
	"crypto/sha256"
	. "dangle/src/tcoloring"
	Tutorials "dangle/src/tutorials"
	"fmt"
	"strings"
)

// 🚧

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

type Chain struct {
	blocks []*Block
}

func main() {
	fmt.Println(strings.Repeat("-", 37), "🔨")
	// execTutorialProgram()
	chain := InitChain()
	chain.AddBlock("1")
	chain.AddBlock("2")
	chain.AddBlock("3")
	for i, block := range chain.blocks {
		fmt.Printf("Block #%v:\nHash: %x\nData: %s\nPrevious hash: %x\n", i, block.Hash, block.Data, block.PreviousHash)
	}
	// TODO -> Making of a 'block-chain' base package.
}

func (chain *Chain) AddBlock(data string) {
	previousBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// func Genesis() *Block {
// 	return CreateBlock("Genesis Block", []byte{})
// }

func InitChain() *Chain {
	return &Chain{[]*Block{CreateBlock("Genesis Block", []byte{})}}
}

// 🚧

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.DeriveHash()
	return block
}

// A 'receiver' - (block *Block) - like a class that is needed for an instance to call some 'class' function.
func (block *Block) DeriveHash() {
	info := bytes.Join([][]byte{block.Data, block.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	block.Hash = hash[:]
}

// 🚧

func execTutorialProgram() {
	Yellow("Note: Some dependencies use deprecated macOS modules and packages!\n")
	CyanBold("Tutorial program: START")
	Tutorials.First()
	Tutorials.Second()
	Tutorials.Third()
	Tutorials.Fourth()
	Tutorials.Fifth()
	Tutorials.Sixth()
	Yellow("Tutorials: FINISHED\n" +
		"Checked out:\n" +
		"* Go's switch statement (fallthrough, type-switch: interface{}.type cases, early break)\n" +
		"* Or (|) condition short circuiting\n" +
		"* Looping (labels, ranges)\n" +
		"* Defer (continue later exec of the program, before main finishes exec, LIFO), panic (Go's exceptions), recover (saving the program)\n" +
		"* Continuing in the Go Playground\n" +
		"* Continuing in the project implementation")
	CyanBold("Tutorial program: END\n")
}
