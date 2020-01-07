package main

// Go looks nice. 🐘

import (
	"fmt"
	. "gotuts/src/tcoloring"
	Tutorials "gotuts/src/tutorials"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("-", 2*37), "🏁")
	defer fmt.Println(strings.Repeat("-", 2*37), "🏁")
	func() {
		execTutorialProgram()
	}()
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
		"* Putting Go on hold, specialize and go deeper SE and CS 👍")
	CyanBold("Tutorial program: END\n")
}
