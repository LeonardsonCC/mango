package manager

import (
	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/mango/scrappers/mangadex"
)

var scrp = map[string]scrappers.Scrapper{
	"MangaDex": mangadex.NewScrapper(),
}

type Manager struct {
	scrappers map[string]scrappers.Scrapper
	output    string
	language  string
}

func NewManager() *Manager {
	return &Manager{
		scrappers: scrp,
		language:  "pt-br",
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

func (m *Manager) SetLanguage(language string) {
	m.language = language
	for _, scrp := range m.scrappers {
		scrp.SetLanguage(language)
	}
}
