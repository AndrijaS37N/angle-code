package tutorials

import (
	"crypto/hmac"
	"crypto/sha256"
	. "dangle/src/tcoloring"
	"encoding/hex"
	"fmt"
	"reflect"
)

const (
	messageToHelena string = "I think of you."
	symbol          string = "<3"
)

type Husky struct {
	id      int // To be visible outside of this package, use uppercase member names.
	name    string
	friends []string
}

type Animal struct {
	name string `String name minimum characters: 2`
	size string
}

type Elephant struct {
	Animal  // Go's struct complementing, Go's "inheritance". Composition via embedding!
	earSize string
	weather  string
}

var my Elephant = Elephant{
	Animal: Animal{
		name: "Andrew",
		size: "Small",
	},
	earSize: "Big",
	weather:  "Stormy",
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
		message  string
		username string
		password string
		symbol   string
	}{username: "Anon K'eyush", password: mySHA, message: "", symbol: symbol}
	anonHusky.message = "Ruf ruf!"
	fmt.Println(anonHusky)
	anotherAnonHusky := anonHusky // To point to the same struct in memory use &anonHusky.
	anotherAnonHusky.username = "Andrew"
	anotherAnonHusky.message = messageToHelena
	fmt.Println(anotherAnonHusky)
	// Change some Elephant{} params.
	my.weather = "🌤"
	fmt.Println("Elephant:", my)
	// Tags.
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Print(field.Tag)
	/*
		Note about maps and structures:
		Maps, like slices and arrays are referenced types.
		Structures are value types, not effecting other structures that are passed to a new struct.
	*/
	WhiteBold("Fifth personal tutorial: END")
}
