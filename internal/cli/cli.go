package cli

import (
	"github.com/LeonardsonCC/mango/internal/app/scrappers"
)

type Cli struct {
	scrapper scrappers.Scrapper
}

func NewCli(s scrappers.Scrapper) *Cli {
	return &Cli{
		scrapper: s,
	}
}
