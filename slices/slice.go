package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 0, 3)
	var input string
	for {
		fmt.Print("Enter an integer (or 'X' to quit): ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}
		if input == "X" {
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}
		s = append(s, num)
		sort.Ints(s)
		fmt.Println("Sorted slice:", s)
	}
}
