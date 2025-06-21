package main

import (
	"fmt"
	"github.com/trolioSFG/blogconfig"
)

func main() {
	fmt.Println("Blog aggregator")
	c := blogconfig.Read()

	c.SetUser("sergio")

	c = blogconfig.Read()
	fmt.Printf("%+v\n", c)
}
