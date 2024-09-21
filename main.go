package main

import (
	"bufio"
	"fmt"
	api "github.com/Dhairya3124/PokeDex/pokeapi"
	"io"
	"os"
)

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
	commands := api.GetCommands()
	for i := 0; ; i++ {
		input := cli.readLine()
		for k, v := range commands {
			if k == input {
				if err := v.Callback(); err != nil {
					fmt.Println(err)
				}
			}
		}

	}

}
func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
