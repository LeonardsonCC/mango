package cli

import (
	"fmt"

	"github.com/LeonardsonCC/mango/internal/cli/colors"
	"github.com/spf13/cobra"
)

func (c *Cli) List() *cobra.Command {
	return &cobra.Command{
		Use:   "list [name] [chapter]",
		Short: "list chapters from manga",
		Args:  cobra.ExactArgs(1),
		Run:   c.list,
	}
}

func (c *Cli) list(cmd *cobra.Command, args []string) {
	name := args[0]

	results, err := c.manager.ListChapters(name)
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("failed to list chapters: %s", err.Error()))
	}

	if len(results) < 1 {
		fmt.Println(colors.Errors.Sprint("no chapters found for this manga"))
		return
	}

	for k, r := range results {
		fmt.Println(colors.Info.Sprint(k))

		if len(r) == 0 {
			fmt.Println(colors.Errors.Sprint("no chapters found"))
		}

		for _, c := range r {
			fmt.Println(c.Title())
		}
	}
}
