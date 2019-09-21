package main

// Go looks nice. 🐘

import (
	. "dangle/src/tcoloring"
	Tutorials "dangle/src/tutorials"
)

func main() {
	CyanBold("Tut program: START")
	Tutorials.First()
	Tutorials.Second()
	Tutorials.Third()
	Tutorials.Fourth()
	Tutorials.Fifth()
	Tutorials.Sixth()
	Yellow("TUTS FINISHED\n" +
		"Checked out:\n" +
		"* Go's switch statement (fallthrough, type-switch: interface{}.type cases, early break)\n" +
		"* Or (|) condition short circuiting\n" +
		"* Looping - Continuing in the Go Playground -\n")
	CyanBold("Tut program: END\n")
}
