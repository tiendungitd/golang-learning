package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	input = strings.ToLower(input) // Convert the input to lowercase for case-insensitive comparison
	if input[0] != 'i' || input[len(input)-1] != 'n' {
		fmt.Println("Not Found!")
		return
	}
	for i := 1; i < len(input)-1; i++ {
		//fmt.Printf("%c\n", input[i])
		if input[i] == 'a' {
			fmt.Println("Found!")
			break
		}
		fmt.Println("Not Found!")
	}
}
