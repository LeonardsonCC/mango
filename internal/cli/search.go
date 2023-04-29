package cli

import (
	"fmt"

	"github.com/LeonardsonCC/mango/internal/cli/colors"
	"github.com/spf13/cobra"
)

func (c *Cli) Search() *cobra.Command {
	return &cobra.Command{
		Use:   "search [name]",
		Short: "searches by manga name",
		Args:  cobra.ExactArgs(1),
		Run:   c.search,
	}
}

func (c *Cli) search(cmd *cobra.Command, args []string) {
	name := args[0]

	results, err := c.scrapper.SearchManga(name)
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("failed to search by %s: %s", name, err.Error()))
	}

	for _, r := range results {
		fmt.Println(r.Title())
	}
}
