package main

// Go looks nice. 🐘

import (
	. "dangle/src/tcoloring"
	Tutorials "dangle/src/tutorials"
)

func main() {
	CyanBold("Program: START")
	Tutorials.First()
	Tutorials.Second()
	Tutorials.Third()
	Tutorials.Fourth()
	Tutorials.Fifth()
	Tutorials.Sixth()
	CyanBold("Program: END\n")
}
