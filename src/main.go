package main

// Go looks nice. 🐘

import (
	. "./tcoloring"
	Tutorial "./tutorial"
)

func main() {
	CyanBold("Program: START")
	Tutorial.First()
	Tutorial.Second()
	Tutorial.Third()
	Tutorial.Fourth()
	Tutorial.Fifth()
	Tutorial.Sixth()
	CyanBold("Program: END\n")
}
