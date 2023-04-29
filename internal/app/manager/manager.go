package manager

import (
	"github.com/LeonardsonCC/mango/internal/app/scrappers"
	"github.com/LeonardsonCC/mango/internal/app/scrappers/mangalivre"
	"github.com/LeonardsonCC/mango/internal/app/scrappers/muitomanga"
)

type Manager struct {
	scrappers map[string]scrappers.Scrapper
}

func NewManager() *Manager {
	return &Manager{
		scrappers: map[string]scrappers.Scrapper{
			"MuitoManga": muitomanga.NewScrapper(),
			"MangaLivre": mangalivre.NewScrapper(),
		},
	}
}
