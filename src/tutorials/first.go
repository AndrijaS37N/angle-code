package tutorials

import (
	"fmt"
	. "gotuts/src/tcoloring"
	"strconv"
)

var (
	integerVar    int
	stringVar     string
	floatVar32    float32
	floatVar64    float64
	complexVar64  complex64
	complexVar128 complex128
)

func First() {
	WhiteBold("First personal tutorial: START")
	integerVar = 1
	stringVar = "Woo woo!"
	floatVar32 = 1.
	floatVar64 = 1.
	complexVar64 = 1
	complexVar128 = 1
	fmt.Printf("%s %T %T %T %T %T %T\n", "Variable types:", integerVar, stringVar, floatVar32, floatVar64, complexVar64, complexVar128)
	fmt.Printf("%s %v %v %v %v %v %v\n", "Variable values:", integerVar, stringVar, floatVar32, floatVar64, complexVar64, complexVar128)
	// Localized variables, interpreted variables.
	a := 5
	b := 2.4
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	a = 3
	b = 6.4
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	a = int(b)
	b = float64(a)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	// ASCII representation of the number 3.
	c := string(51)
	fmt.Printf("%s %T\n", c, c)
	c = strconv.Itoa(a)
	fmt.Printf("%s %T", c, c)
	WhiteBold("First personal tutorial: END")
}
