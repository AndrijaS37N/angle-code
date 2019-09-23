package tutorials

import (
	. "dangle/src/tcoloring"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"os"
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
	fmt.Print("Enter text: ")
	var text string
	t, err := fmt.Scanln(&text)
	ioFunc(t, err, &text)
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
	// Check if text entered by the user is bigger than 7 in char len.
	fmt.Print("Checking the length of the entered text (text > 7): ", checkLen(&text))
	WhiteBold("Sixth personal tutorial: END")
}

func checkLen(text *string) bool {
	if len(*text) > 7 {
		return true
	} else {
		return false
	}
}

func ioFunc(t int, err error, text *string) string {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return err.Error()
	}
	fmt.Println("Adding '_37' to the end of the string:")
	*text = *text + "_37"
	fmt.Println(t)
	fmt.Println(err)
	fmt.Println("Text:", *text)
	return *text
}
