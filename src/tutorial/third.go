package tutorial

import (
	. "../tcoloring"
	"fmt"
	"math"
)

const (
	integerConstant int     = 7
	floatConstant32 float32 = math.Pi
	floatConstant64 float64 = math.E
	stringConstant  string  = "Waf waf!"
)

const (
	// Iota is scoped to this constant block.
	_ = iota + 3
	iotaOne
	iotaTwo
	iotaThree
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

const (
	admin = 1 << iota
	manager
	user
	// Authorizations.
	seeEU
	seeUSA
	seeAfrica
	seeAsia
)

var roles byte = admin | seeAsia | seeEU

func Third() {
	WhiteBold("Third personal tutorial: START")
	fmt.Printf("%v %T\n", integerConstant, integerConstant)
	fmt.Printf("%v %T\n", floatConstant32, floatConstant32)
	fmt.Printf("%v %T\n", floatConstant64, floatConstant64)
	fmt.Printf("%v %T\n", stringConstant, stringConstant)
	const myConst int = 5
	var myVar int = 4
	fmt.Printf("Constant + variable = variable, so we have: %v %T\n", myConst+myVar, myConst+myVar)
	fmt.Printf("%v %T\n", iotaOne, iotaOne)              // _ (0), 1+3
	fmt.Printf("%v %T\n", iotaTwo, iotaTwo)              // 2+3
	fmt.Printf("%v %T\n", iotaThree, iotaThree)          // 3+3
	fileSize := 5000000.                                 // 5,000,000.00 bytes
	fmt.Printf("%.5f GB %T\n", fileSize/GB, fileSize/GB) // ~ 0.005 GB
	fmt.Printf("GB const iota value: %v %T\n", GB, GB)   // 1,073,741,824 integer
	fmt.Printf("Roles set: admin | seeAsia | seeEU\nBits set: %b\n", roles)
	fmt.Printf("Admin roles? %v\n", admin&roles)
	fmt.Printf("Manager roles? %v", manager&roles)
	WhiteBold("Third personal tutorial: END")
}
