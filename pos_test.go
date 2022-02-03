package main

import "fmt"

func ExampleString_Pos() {
	pos := Pos{0, 2}
	fmt.Println(pos)

	// Output:
	// A2
}

func ExampleConvert() {
	s := "A2"
	pos, _ := Convert(s)
	fmt.Println(pos)

	// Output:
	// A2
}
