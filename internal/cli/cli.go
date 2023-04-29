package cli

import (
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
