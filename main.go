package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	pokecache "github.com/Dhairya3124/PokeDex/pokeCache"
)

type CLI struct {
	out io.Writer
	in  *bufio.Scanner
}
type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

func NewCLI(in io.Reader, out io.Writer) *CLI {
	return &CLI{
		in:  bufio.NewScanner(in),
		out: out,
	}
}

func main() {
	cli := NewCLI(os.Stdin, os.Stdout)
	commands := GetCommands()
	config := new(Config)
	config.Cache = pokecache.NewCache(time.Duration(60) * time.Second)
	for i := 0; ; i++ {
		input := strings.ToLower(cli.readLine())
		for k, v := range commands {
			if k == input {
				if err := v.Callback(config); err != nil {
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
