package cli

import (
	"bytes"
	"fmt"

	"github.com/LeonardsonCC/mango/internal/cli/colors"
	"github.com/LeonardsonCC/mango/internal/cli/spinner"
	"github.com/LeonardsonCC/mango/mango/scrappers"
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

	loading := make(chan struct{})

	var (
		results map[string][]*scrappers.SearchMangaResult
		errs    map[string]error
	)

	go func() {
		results, errs = c.manager.Search(name)
		loading <- struct{}{}
	}()

	spinner.Loading(loading)

	for k, scrapper := range results {
		fmt.Print(c.genOutput(k, scrapper, errs[k]))
	}
}

func (c *Cli) genOutput(name string, results []*scrappers.SearchMangaResult, err error) string {
	var str bytes.Buffer
	str.WriteString(colors.Info.Sprintf("%s\n", name))

	if err != nil {
		str.WriteString(colors.Errors.Sprintf("Failed to search: %v\n", err))
		str.WriteString("\n\n")
		return str.String()
	}

	if len(results) == 0 {
		str.WriteString(colors.Warning.Sprint("no results...\n"))
	}

	for _, r := range results {
		str.WriteString(r.Title() + "\n")
	}

	str.WriteString("\n\n")
	return str.String()
}
