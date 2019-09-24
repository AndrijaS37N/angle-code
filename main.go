package main

// Go looks nice. 🐘

import (
	Protocol "dangle/src/blockchain"
	. "dangle/src/tcoloring"
	Tutorials "dangle/src/tutorials"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// execTutorialProgram()
	fmt.Println(strings.Repeat("-", 2*37), "🏁")
	defer fmt.Println(strings.Repeat("-", 2*37), "🏁")
	func() {
		chain := Protocol.InitChain()
		for i := 1; i <= 7; i++ {
			chain.AddBlock(strconv.Itoa(i))
		}
		fmt.Println()
		for i, block := range chain.Blocks {
			fmt.Printf("Block #%v:\nPrevious hash: %x\nData: %s\nHash: %x\n", i, block.PreviousHash, block.Data, block.Hash)
			pow := Protocol.NewProof(block)
			fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
			fmt.Println()
		}
		Yellow("WIP")
		fmt.Println()
	}()
	// TODO -> Making of a 'block-chain' base package.
}

// 🧩👇
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
