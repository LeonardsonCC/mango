package cli

import (
	"fmt"
	"os"

	"github.com/LeonardsonCC/mango/internal/cli/colors"
	"github.com/spf13/cobra"
)

func (c *Cli) Download() *cobra.Command {
	return &cobra.Command{
		Use:   "download [name] [chapter]",
		Short: "download the manga in pdf",
		Args:  cobra.ExactArgs(2),
		Run:   c.download,
	}
}

func (c *Cli) download(cmd *cobra.Command, args []string) {
	name := args[0]
	chapter := args[1]

	fmt.Printf("searching %s...\n", name)
	r, err := c.scrapper.SearchManga(name)
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("failed to find: %s", name))
	}

	fmt.Printf("searching chapter %s...\n", chapter)
	chap, err := c.scrapper.SearchChapter(r[0].Url(), chapter)
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("failed to find chapter: %s", chapter))
	}

	if len(chap) < 1 {
		fmt.Println(colors.Errors.Sprintf("chapter not found: %s", chapter))
		return
	}

	fmt.Printf("downloading chapter %s from %s\n", chap[0].Title(), r[0].Title())
	manga, err := c.scrapper.Download(chap[0].Url())
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("failed to download chapter: %s", chapter))
	}

	filename := fmt.Sprintf("./%s.pdf", manga.Title)
	f, _ := os.Create(filename)
	defer f.Close()

	_, err = f.Write(manga.Buffer.Bytes())
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("failed to save pdf"))
	}

}
