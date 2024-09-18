package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CLICommand struct {
	name        string
	description string
	callback    func() error
}

func NewCLICommand() map[string]CLICommand {
	return map[string]CLICommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
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
	commands := NewCLICommand()
	for _, v := range commands {
		fmt.Printf("%s : %s\n", v.name, v.description)
	}

	return nil
}

type CLI struct {
	out io.Writer
	in  *bufio.Scanner
}

func NewCLI(in io.Reader, out io.Writer) *CLI {
	return &CLI{
		in:  bufio.NewScanner(in),
		out: out,
	}
}

func main() {
	cli := NewCLI(os.Stdin, os.Stdout)
	commands := NewCLICommand()
	for i := 0; ; i++ {
		input := cli.readLine()
		for k, v := range commands {
			if k == input {
				v.callback()
			}
		}

	}

}
func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
