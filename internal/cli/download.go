package cli

import (
	"fmt"
	"os"

	"github.com/LeonardsonCC/mango/internal/cli/colors"
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

	err := c.manager.DownloadChapter(name, chapter)
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("error: %v", err))
		os.Exit(1)
	}
}

func (c *Cli) downloadManga(cmd *cobra.Command, args []string) {
	name := args[0]

	err := c.manager.DownloadManga(name)
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("error: %v", err))
		os.Exit(1)
	}
}
