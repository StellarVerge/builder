package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "builder"
	app.Usage = "Build Stellar Verge"
	app.Version = "0.0.1"
	app.Run(os.Args)
}
