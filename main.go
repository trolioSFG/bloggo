package main

import (
	"fmt"
	"github.com/trolioSFG/blogconfig"
	"os"
)

type state struct {
	cfg *blogconfig.Config
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bloggo <command> [<cmd_args>...]")
		os.Exit(1)
	}

	fmt.Println("Blog aggregator")
	c := blogconfig.Read()
	s := state{}
	s.cfg = &c
/**
	c.SetUser("sergio")

	c = blogconfig.Read()
	fmt.Printf("%+v\n", c)

	cmd := command{name: "login", args: []string{"sergio"}}
	// c.name = "login"
	// c.args := []string{ "sergio" }
	err := handlerLogin(&s, cmd)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", s.cfg)
	}

*/

	cmd := command {
		name: os.Args[1],
		args: os.Args[2:],
	}

	cmds := commands{}
	cmds.cmds = make(map[string]func(*state, command) error, 0)

	cmds.register("login", handlerLogin)
	err := cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
