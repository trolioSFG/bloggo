package main

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	"github.com/trolioSFG/blogconfig"
	"github.com/trolioSFG/database"
	_ "github.com/lib/pq"

	"context"
)

type state struct {
	cfg *blogconfig.Config
	db *database.Queries
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}


func main() {

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


	cmds := commands{}
	cmds.cmds = make(map[string]func(*state, command) error, 0)

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerAddFeedFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	// TEST
	// cmds.register("scrape", scrapeFeeds)

	if len(os.Args) < 2 {
		fmt.Println("Usage: bloggo <command> [<cmd_args>...]")
		fmt.Println("Available commands:")
		for name, _ := range cmds.cmds {
			fmt.Println(name)
		}
		os.Exit(1)
	}

	cmd := command {
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
