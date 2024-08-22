package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	name    string
	address string
}

func main() {
	// Prompt the user to enter a name
	p := Person{}
	fmt.Print("Enter a name: ")
	fmt.Scanln(&p.name)

	// Prompt the user to enter an address
	fmt.Print("Enter an address: ")
	fmt.Scanln(&p.address)

	// Create a map to store the name and address
	data := map[string]string{
		"name":    p.name,
		"address": p.address,
	}

	// Convert the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print the JSON object
	fmt.Println(string(jsonData))
}
