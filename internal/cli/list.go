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

	results, err := c.scrapper.SearchManga(name)
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("failed to search by %s: %s", name, err.Error()))
	}

	if len(results) < 1 {
		fmt.Println(colors.Errors.Sprint("manga not found"))
		return
	}

	m := results[0]
	chapters, err := c.scrapper.SearchChapter(m.Url(), "")
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("failed to search chapters: %s", err.Error()))
	}

	if len(chapters) < 1 {
		fmt.Println(colors.Errors.Sprint("no chapters found for this manga"))
		return
	}

	for _, r := range chapters {
		fmt.Println(r.Title())
	}
}
