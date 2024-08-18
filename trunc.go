package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	//s2 := arr[2:5]
	fmt.Printf("%v and cap %v", len(s1), cap(s1))
}
