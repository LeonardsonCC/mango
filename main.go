package main

import (
	"os"

	"github.com/LeonardsonCC/mango/cmd"
)

func main() {
	if len(os.Args) > 1 {
		cmd.NewCli().Start()
	} else {
		cmd.NewTui().Start()
	}
}
