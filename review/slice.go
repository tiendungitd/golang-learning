package main

import (
	"fmt"
)

func main() {
	//s := make([]int, 0, 3)
	s := []int{1, 2, 3}
	s = append(s, 100)
	s = append(s, 300)
	fmt.Println(len(s), cap(s))
}
