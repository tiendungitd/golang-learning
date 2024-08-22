package main

import (
	"fmt"

	"github.com/victorops/go-victorops/victorops"
)

func main() {

	// Client initialization
	victoropsClient := victorops.NewClient(apiID, apiKey, "https://api.victorops.com")

	// Get all users in an account
	userList, _, err := victoropsClient.GetAllUsers()
	if err != nil {
		panic(err)
	}

	// Create a new victorops team
	team := victorops.Team{
		Name: "Test Team",
	}
	newTeam, details, err := victoropsClient.CreateTeam(&team)
	if err != nil {
		panic(err)
	}

	if details.StatusCode != 200 {
		panic(fmt.Errorf("failed to create team (%d): %s", details.StatusCode, details.ResponseBody))
	}

	fmt.Printf("Created team: %s\n", newTeam.Name)
}
