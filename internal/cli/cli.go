package cli

import (
	"os"

	"github.com/LeonardsonCC/mango/internal/app/manager"
)

type Cli struct {
	manager *manager.Manager
}

func NewCli() *Cli {
	return &Cli{
		manager: manager.NewManager(),
	}
}
func (c *Cli) SetScrapper(scrapper string) {
	c.manager.SetScrapper(scrapper)
}

func (c *Cli) SetOutput(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	c.manager.SetOutput(path)

	return nil
}
