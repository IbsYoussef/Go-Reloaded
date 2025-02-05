package main

import (
	"fmt"
	"go-reloaded/internal/textmod"
)

func main() {
	test := "one (up)"
	test2 := "one, two, three, four (up, 4)"
	test3 := "ONE TWO THREE FOUR (low, 4)"
	test4 := "one two three four (cap, 4)"
	test5 := "TWO (low)"
	test6 := "three (cap)"
	fmt.Println(textmod.ChangeCaseFinal(test))
	fmt.Println(textmod.ChangeCaseFinal(test2))
	fmt.Println(textmod.ChangeCaseFinal(test3))
	fmt.Println(textmod.ChangeCaseFinal(test4))
	fmt.Println(textmod.ChangeCaseFinal(test5))
	fmt.Println(textmod.ChangeCaseFinal(test6))
}
