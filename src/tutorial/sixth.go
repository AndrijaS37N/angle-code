package tutorial

import (
	. "../tcoloring"
	"fmt"
)

var numbers = map[int]string{
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var pop int

func Sixth() {
	WhiteBold("Sixth personal tutorial: START")
	if pop, ok := numbers[4]; ok { // Boolean 'ok' is evaluated. Int pop is the searched for value.
		fmt.Println(pop)
	} else {
		fmt.Println("Element not found.") // Boolean 'ok' is false.
	}
	if pop >= 3 {
		fmt.Println("Pop is greater than 3.")
	}

	// TODO -> WIP

	Yellow("WIP")
	WhiteBold("Sixth personal tutorial: END")
}
