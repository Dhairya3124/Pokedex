package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"time"

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
		"catch": {
			Name:        "catch",
			Description: "Attempts to catch the given pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Shows the details of the pokemon which is already caught",
			Callback:    commandInspect,
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
	Config.CurrentAreaPokemon = []string{}
	for _, encounter := range locationDetailedResponse.PokemonEncounters {
		name := encounter.Pokemon.Name
		Config.CurrentAreaPokemon = append(Config.CurrentAreaPokemon, name)

		fmt.Printf(" - %v\n", name)
	}
	Config.CurrentArea = params
	return nil
}
func commandCatch(Config *Config,params string)error{
	pokemonToCatch:=params
	fmt.Println(pokemonToCatch)
	
	if Config.CurrentArea == "" {
		fmt.Println("Can't catch anything til you go somewhere.")
		return nil
	}

	if !slices.Contains(Config.CurrentAreaPokemon, pokemonToCatch) {
		fmt.Println("Pokemon is not found in this area.")
		return nil
	}
	if _, ok := Config.Pokedex[pokemonToCatch]; ok {
		fmt.Println("You already have this pokemon")
		return nil
	}

	
	resp, err := pokeapi.FetchPokemonDetailsAPI(pokemonToCatch, Config.Cache)
	if err != nil {
		return err
	}
	// fmt.Println(resp)
	randCatchProb := rand.Float32()
	experienceMultiplier := 1 - resp.BaseExperience/1000

	catchProb := randCatchProb * float32(experienceMultiplier)

	fmt.Printf("Throwing pokeball at %v...\n", pokemonToCatch)

	time.Sleep(500 * time.Millisecond)

	if catchProb >= 0.5 {
		fmt.Printf("%v was caught!\n", pokemonToCatch)
		Config.Pokedex[pokemonToCatch] = *resp
	} else {
		fmt.Printf("%v escaped!\n", pokemonToCatch)
	}
	return nil
}
func commandInspect(Config *Config,params string)error{
	pokemonToInspect:=params
	if pokemonResp,ok:=Config.Pokedex[pokemonToInspect];!ok{
		fmt.Println("Pokemon not found")
	}else{
		fmt.Printf("Name: %v\n",pokemonResp.Name)
		fmt.Printf("Height: %v\n",pokemonResp.Height)
		fmt.Printf("Weight: %v\n",pokemonResp.Weight)
		fmt.Println("Stats:")
		for _,stat:= range pokemonResp.Stats{
			fmt.Printf("- %v:%v\n",stat.Stat.Name,stat.BaseStat)
		}
		fmt.Println("Types:")
		for _,typ:= range pokemonResp.Types{
			fmt.Printf("- %v\n",typ.Type.Name)
		}
		
	}
	
	
	return nil
}