package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Dhairya3124/PokeDex/pokeapi"
)

type CLICommand struct {
	Name        string
	Description string
	Callback    func(Config *Config, params string) error
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
		"explore":{

			Name:        "explore",
			Description: "Show detailed response of location",
			Callback:    showLocationExplore,			
		},
	}
}
func commandExit(Config *Config, params string) error {
	os.Exit(0)
	return nil
}
func commandHelp(Config *Config, params string) error {
	fmt.Println("Welcome to PokeDex!")
	fmt.Println("Usage:")
	commands := GetCommands()
	for _, v := range commands {
		fmt.Printf("%s : %s\n", v.Name, v.Description)
	}

	return nil
}

func showResponseofAPI(Config *Config, params string) error {
	locationResponse, err := pokeapi.FetchPokeAPI(Config.Next, Config.Cache)
	if err != nil {
		return err
	}
	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	Config.Next = locationResponse.Next
	Config.Previous = locationResponse.Previous
	return nil

}
func showPrevResponseofAPI(Config *Config, params string) error {
	if Config.Previous == "" {
		return errors.New("no previous URL found")
	}
	locationResponse, err := pokeapi.FetchPokeAPI(Config.Previous, Config.Cache)
	if err != nil {
		return err
	}

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	Config.Next = locationResponse.Next
	Config.Previous = locationResponse.Previous
	return nil

}
func showLocationExplore(Config *Config, params string)error{
	locationDetailedResponse,err:=pokeapi.FetchPokeExploreAPI(params,Config.Cache)
	if err!=nil{
		return err
	}
	for _, encounter := range locationDetailedResponse.PokemonEncounters {
		name := encounter.Pokemon.Name
		fmt.Printf(" - %v\n", name)
	}
	return nil
}