package pokeapi

import (
	"fmt"
	"os"
)

type CLICommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]CLICommand {
	return map[string]CLICommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Show location of the maps",
			Callback:    showResponseofAPI,
		},
		"mapb": {
			Name:        "map",
			Description: "Show location of the maps",
			Callback:    showPrevResponseofAPI,
		},
	}
}
func commandExit() error {
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("Welcome to PokeDex!")
	fmt.Println("Usage:")
	commands := GetCommands()
	for _, v := range commands {
		fmt.Printf("%s : %s\n", v.Name, v.Description)
	}

	return nil
}

var pokeLocationURL = "https://pokeapi.co/api/v2/location-area/"
var count = 0

func showResponseofAPI() error {
	locationResponse, finalCount, err := FetchPokeAPI(pokeLocationURL, count)
	if err != nil {
		return err
	}
	fmt.Println(locationResponse)
	fmt.Println(finalCount)
	count = finalCount
	pokeLocationURL = locationResponse.Next
	return nil

}
func showPrevResponseofAPI() error {
	locationResponse, finalCount, err := FetchPokeAPI(pokeLocationURL, count)
	if err != nil {
		return err
	}
	if locationResponse.Previous != nil {
		locationResponse, finalCount, err = FetchPokeAPI(*locationResponse.Previous, count)
		if err != nil {
			return err
		}
		fmt.Println(locationResponse)

	} else {
		fmt.Println("No previous URL found")
	}
	count = finalCount
	if locationResponse.Previous != nil {
		pokeLocationURL = *locationResponse.Previous
	}
	return nil

}
