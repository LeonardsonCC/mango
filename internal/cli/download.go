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

	err := c.manager.Download(name, chapter)
	if err != nil {
		fmt.Println(colors.Errors.Sprintf("error: %v", err))
		os.Exit(1)
	}
}
