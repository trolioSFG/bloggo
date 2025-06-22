package main

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	"github.com/trolioSFG/blogconfig"
	"github.com/trolioSFG/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *blogconfig.Config
	db *database.Queries
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

	db, err := sql.Open("postgres", c.DbURL)
	if err != nil {
		log.Fatalf("Could not connect to postgresql: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	s.db = dbQueries


	cmd := command {
		name: os.Args[1],
		args: os.Args[2:],
	}

	cmds := commands{}
	cmds.cmds = make(map[string]func(*state, command) error, 0)

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
