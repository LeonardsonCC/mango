package cli

import (
	"bytes"
	"fmt"

	"github.com/LeonardsonCC/mango/internal/cli/colors"
	"github.com/LeonardsonCC/mango/internal/cli/spinner"
	"github.com/LeonardsonCC/mango/mango/scrappers"
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

	loading := make(chan struct{})

	var (
		results map[string][]*scrappers.SearchChapterResult
		errs    map[string]error
	)

	go func() {
		results, errs = c.manager.ListChapters(name)
		loading <- struct{}{}
	}()

	spinner.Loading(loading, "Listing chapters...")

	for k, r := range results {
		fmt.Print(c.genOutputList(k, r, errs[k]))
	}
}

func (c *Cli) genOutputList(name string, results []*scrappers.SearchChapterResult, err error) string {
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
		str.WriteString(r.Chapter() + " - " + r.Title() + "\n")
	}

	str.WriteString("\n\n")
	return str.String()
}
