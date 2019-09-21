package tutorial

import (
	. "dangle/src/tcoloring"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var numbers = map[int]string{
	10: "X",
	9:  "IX",
	5:  "V",
	4:  "IV",
	1:  "I",
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
	// Trying out the cross-platform GUI library of Go.
	fmt.Println("Opening a GUI window")
	myApp := app.New()
	w := myApp.NewWindow("Hello From Go")
	w.CenterOnScreen()
	fmt.Println(".....")
	w.Resize(fyne.NewSize(700, 300))
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello World!"),
		widget.NewButton("Quit", func() {
			myApp.Quit()
		}),
	))
	w.ShowAndRun()
	fmt.Println("GUI window closed")

	// TODO -> WIP

	Yellow("WIP")
	WhiteBold("Sixth personal tutorial: END")
}
