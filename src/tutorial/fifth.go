package tutorial

import (
	. "../tcoloring"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	messageToHelena string = "I think of you."
	symbol string = "❤️"
)

type Husky struct {
	id      int // To be visible outside of this package, use uppercase member names.
	name    string
	friends []string
}

func Fifth() {
	WhiteBold("Fifth personal tutorial: START")
	myMap := map[string]int{
		"Aaa": 2412,
		"Bee": 23,
		"Eee": 612535,
		"Ree": 513,
	}
	fmt.Println(myMap)
	// Memory address printing.
	print("Memory address: ", myMap, "\n")
	// Slices, functions and maps cannot be used for equivalence checking.
	m := map[[3]int]string{}
	fmt.Println(myMap, m)
	n := make(map[int]int)
	n = map[int]int{
		1: 1000,
		4: 200000,
		6: 214121,
	}
	fmt.Println(n)
	fmt.Println("The value of key '4' in map n:", n[4])
	n[7] = 5
	/*
		Unknown in what order will this element be inserted
		into the map, unlike for arrays and slices.
	*/
	fmt.Println("The value of key '7' in map n:", n[7])
	fmt.Println(n)
	delete(n, 7)
	fmt.Println(n)
	// Checking the presence of key '9'.
	_, ok := n[9]
	fmt.Println("Is key '9' present in map n:", ok)
	// These reference the same underlying structure.
	j := n
	delete(j, 4)                       // That's why when I do this...
	fmt.Println("Maps n and j:", n, j) // ...This happens!
	myFriend := Husky{
		id:   1,
		name: "K'eyush The Stunt Dog",
		friends: []string{
			"Helena",
			"Eugen",
			"Andrew",
		},
	}
	/*
		Best to state all members as they are declared
		(in case of changing/adding/deleting struct members).
	*/
	fmt.Println(myFriend)
	fmt.Println("Friends of", myFriend.name, "are", myFriend.friends)
	fmt.Println("First friend of", myFriend.name, "is", myFriend.friends[0])
	secret := "WAF WAF RECIPE"
	data := "Description: Need lots of wuf wuf, love, a couple of barks, etc."
	myHMAC := hmac.New(sha256.New, []byte(secret))
	myHMAC.Write([]byte(data))
	mySHA := hex.EncodeToString(myHMAC.Sum(nil))
	anonHusky := struct {
		message string
		username string
		password string
		symbol string
	}{username: "Anon K'eyush", password: mySHA, message: "", symbol: symbol}
	anonHusky.message = "Ruf ruf!"
	fmt.Println(anonHusky)
	anotherAnonHusky := anonHusky // To point to the same struct in memory use &anonHusky.
	anotherAnonHusky.username = "Andrew"
	anotherAnonHusky.message = messageToHelena
	fmt.Println(anotherAnonHusky)

	// TODO -> WIP Go's "Inheritance"

	Yellow("WIP")
	WhiteBold("Fifth personal tutorial: END")
}
