package cmd

import (
	"fmt"
	"os"

	"github.com/LeonardsonCC/mango/internal/cli"
	"github.com/spf13/cobra"
)

type Cli struct{}

func NewCli() *Cli {
	return &Cli{}
}

func (*Cli) Start() {
	c := cli.NewCli()

	root := cobra.Command{}
	root.AddCommand(c.Download())
	root.AddCommand(c.Search())
	root.AddCommand(c.List())

	if err := root.Execute(); err != nil {
		fmt.Println("failed to run command: ", err)
		os.Exit(1)
	}
}
