package cli

import (
	"fmt"
	"os"

	"github.com/LeonardsonCC/mango/internal/cli/colors"
	"github.com/LeonardsonCC/mango/internal/cli/spinner"
	"github.com/spf13/cobra"
)

func (c *Cli) Download() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download [name] [chapter]",
		Short: "download the manga in pdf",
		Args:  cobra.ExactArgs(2),
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:   "chapter [name] [chapter]",
			Short: "download the chapter from a manga in pdf",
			Args:  cobra.ExactArgs(2),
			Run:   c.downloadChapter,
		})

	cmd.AddCommand(
		&cobra.Command{
			Use:   "manga [name]",
			Short: "download the chapter from a manga in pdf",
			Args:  cobra.ExactArgs(1),
			Run:   c.downloadManga,
		})

	return cmd
}

func (c *Cli) downloadChapter(cmd *cobra.Command, args []string) {
	name := args[0]
	chapter := args[1]

	loading := make(chan struct{})

	var err error
	go func() {
		err = c.manager.DownloadChapter(name, chapter)
		loading <- struct{}{}
	}()

	spinner.Loading(loading, "Downloading chapter...")

	if err != nil {
		fmt.Println(colors.Errors.Sprintf("error: %v", err))
		os.Exit(1)
	}
	fmt.Println(colors.Info.Sprintf("Downloaded with success"))
}

func (c *Cli) downloadManga(cmd *cobra.Command, args []string) {
	name := args[0]

	loading := make(chan struct{})

	var err error
	go func() {
		err = c.manager.DownloadManga(name)
		loading <- struct{}{}
	}()

	spinner.Loading(loading, "Downloading all chapters...")

	if err != nil {
		fmt.Println(colors.Errors.Sprintf("error: %v", err))
		os.Exit(1)
	}
	fmt.Println(colors.Info.Sprintf("Downloaded with success"))
}
