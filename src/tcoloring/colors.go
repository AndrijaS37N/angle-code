package tcoloring

import "fmt"

func CyanBold(text string) {
	fmt.Printf("\033[1;36m%s\033[0m", text)
}

func WhiteBold(text string) {
	fmt.Printf("\n\033[1;30m%s\033[0m\n", text)
}

func Yellow(text string) {
	fmt.Print("\033[0;33m", text, "\033[0m\b")
}