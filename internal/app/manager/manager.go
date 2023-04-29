package manager

import (
	"github.com/LeonardsonCC/mango/internal/app/scrappers"
	"github.com/LeonardsonCC/mango/internal/app/scrappers/mangalivre"
	"github.com/LeonardsonCC/mango/internal/app/scrappers/muitomanga"
)

var scrp = map[string]scrappers.Scrapper{
	"MuitoManga": muitomanga.NewScrapper(),
	"MangaLivre": mangalivre.NewScrapper(),
}

type Manager struct {
	scrappers map[string]scrappers.Scrapper
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
