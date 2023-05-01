package manager

import (
	"github.com/LeonardsonCC/mango/internal/mango/scrappers"
	"github.com/LeonardsonCC/mango/internal/mango/scrappers/mangalivre"
	"github.com/LeonardsonCC/mango/internal/mango/scrappers/muitomanga"
)

var scrp = map[string]scrappers.Scrapper{
	"MuitoManga": muitomanga.NewScrapper(),
	"MangaLivre": mangalivre.NewScrapper(),
}

type Manager struct {
	scrappers map[string]scrappers.Scrapper
	output    string
}

func NewManager() *Manager {
	return &Manager{
		scrappers: scrp,
	}
}

func (m *Manager) SetScrapper(scrapper string) {
	s := scrp
	if scrapper != "" {
		if v, found := scrp[scrapper]; found {
			s = map[string]scrappers.Scrapper{
				scrapper: v,
			}
		}
	}
	m.scrappers = s
}

func (m *Manager) SetOutput(output string) {
	m.output = output
}
