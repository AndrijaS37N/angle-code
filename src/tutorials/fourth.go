package tutorials

import (
	. "dangle/src/tcoloring"
	"fmt"
)

func Fourth() {
	WhiteBold("Fourth personal tutorial: START")
	grades := [...]int{1, 2, 3}
	fmt.Printf("Grades: %v\n", grades)
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "K'eyush"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("First student: %v\n", students[0])
	fmt.Println("Number of students:", len(students))
	arrayOne := [...]int{1, 2, 3}
	arrayTwo := arrayOne // This is a copy.
	arrayTwo[1] = 7
	fmt.Println(arrayOne)
	fmt.Println(arrayTwo)
	newArrayOne := [...]int{1, 2, 3}
	newArrayTwo := &newArrayOne // Connected by pointers. Same memory.
	newArrayTwo[1] = 7
	fmt.Println(newArrayOne)
	fmt.Println(newArrayTwo)
	// Slices are referenced types. Pointing to the same values in the underlying array.
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	makeVarSlice := make([]int, 3, 100)
	/*
		100 is the capacity.
		The slice length is 3, but the underlying array
		has a length of 100 elements.
	*/
	fmt.Printf("Slice: %v %T\n", makeVarSlice, makeVarSlice)
	fmt.Println("Length", len(makeVarSlice))
	fmt.Println("Capacity", cap(makeVarSlice))
	slice = append(slice, 7, 3, 252, 325, 23, 56)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, []int{1, 2, 3, 4, 5, 6, 7}...) // Spread '...' operator. Spread the new slice out into individual arguments.
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Print(cap(slice))
	WhiteBold("Fourth personal tutorial: END")
}
