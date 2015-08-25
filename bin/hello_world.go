package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Hello World"
	app.Usage = "Say hello to the world!"
	app.Action = func(c *cli.Context) {
		println("Hello world!")
	}

	app.Run(os.Args)
}
