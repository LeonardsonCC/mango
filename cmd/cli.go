package cmd

import (
	"fmt"

	"github.com/LeonardsonCC/mango/internal/app/scrappers/mangalivre"
	"github.com/LeonardsonCC/mango/internal/cli"
	"github.com/spf13/cobra"
)

type Cli struct{}

func NewCli() *Cli {
	return &Cli{}
}

func (*Cli) Start() {
	sc := mangalivre.NewScrapper()
	c := cli.NewCli(sc)

	root := cobra.Command{}
	root.AddCommand(c.Download())
	root.AddCommand(c.Search())
	root.AddCommand(c.List())

	if err := root.Execute(); err != nil {
		fmt.Println("failed to run command: ", err)
	}
}
