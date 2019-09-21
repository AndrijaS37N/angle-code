package tutorial

import (
	. "dangle/src/tcoloring"
	"fmt"
)

var (
	// Booleans are their own type (not aliases for integers).
	n = false
	m bool
)

func Second() {
	WhiteBold("Second personal tutorial: START")
	n = true
	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", m, m)
	// Bad name kind of.
	var unsignedInt64Variable uint64 = 55
	fmt.Printf("%v %T\n", unsignedInt64Variable, unsignedInt64Variable)
	fmt.Printf("%v %T\n", unsignedInt64Variable+uint64(7), unsignedInt64Variable)
	binaryAce := 10                    // 1010
	binaryBee := 5                     // 0101
	fmt.Println(binaryAce & binaryBee) // 0000
	fmt.Println(binaryAce | binaryBee) // 1111
	// XOR operation.
	fmt.Println(binaryAce ^ binaryBee) // 1111
	// AND NOT (not to be misunderstood with NOT AND) operation.
	fmt.Println(binaryAce & ^binaryBee) // 1010
	a := 8                              // 2^3
	fmt.Println(a << 3)                 // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3)                 // 2^3 / 2^3 = 2^0
	// Decimal numbers (floating point-64 number).
	number := 3.14
	number = 15.4e103
	number = 5.2e-35
	fmt.Printf("%v %T\n", number, number)
	fmt.Println("The complex number is:", complexVar64)
	fmt.Printf("Real part: %v %T\n", real(complexVar64), real(complexVar64))
	fmt.Printf("Imaginary part: %v %T\n", imag(complexVar64), imag(complexVar64))
	// Strings are any UTF-8 characters.
	text := "POOS "
	fmt.Printf("%v %T\n", string(text[1]), text[1])
	addedText := "& WOOS"
	fmt.Printf("%v %T\n", text+addedText, text+addedText)
	bytesVariable := []byte(text)
	fmt.Printf("%v %T\n", bytesVariable, bytesVariable)
	// Runes are any UTF-32 characters. Aliases for integer-32 numbers.
	runeVariable := 'r'
	fmt.Printf("%v %T", runeVariable, runeVariable)
	WhiteBold("Second personal tutorial: END")
}
