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

var (
	c        = cli.NewCli()
	scrapper = ""
	output   = ""
)

func (*Cli) Start() {
	cobra.OnInitialize(initialize)

	root := cobra.Command{}
	root.PersistentFlags().StringVarP(&scrapper, "site", "s", "", "Specify the one scrapper, supported: MuitoManga, MangaLivre")
	root.PersistentFlags().StringVarP(&output, "output", "o", "", "Output folder")

	root.AddCommand(c.Download())
	root.AddCommand(c.Search())
	root.AddCommand(c.List())

	if err := root.Execute(); err != nil {
		fmt.Println("failed to run command: ", err)
		os.Exit(1)
	}
}

func initialize() {
	c.SetScrapper(scrapper)

	if output != "" {
		if err := c.SetOutput(output); err != nil {
			fmt.Printf("failed to set output folder %s\n", output)
		}
	}
}
