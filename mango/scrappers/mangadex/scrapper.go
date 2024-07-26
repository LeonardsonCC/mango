package mangadex

import (
	"fmt"

	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/gocolly/colly"
)

const ScrapperName = "MangaDex"

var mainURL = "mangadex.org"
var apiMainURL = "api.mangadex.org"
var imagesURL = []string{}

type Scrapper struct {
	name           string
	apiBaseURL     string
	baseURL        string
	imagesBaseURLs []string
	language       string
	Colly          *colly.Collector
}

func NewScrapper() scrappers.Scrapper {
	c := colly.NewCollector(
		colly.AllowedDomains(mainURL),
		colly.AllowURLRevisit(),
	)

	imgsBaseUrls := make([]string, len(imagesURL))
	for i, url := range imagesURL {
		imgsBaseUrls[i] = fmt.Sprintf("https://%s", url)
	}

	return &Scrapper{
		name:           ScrapperName,
		Colly:          c,
		baseURL:        fmt.Sprintf("https://%s", mainURL),
		apiBaseURL:     fmt.Sprintf("https://%s", apiMainURL),
		imagesBaseURLs: imgsBaseUrls,
		language:       "pt-br",
	}
}

func (s *Scrapper) Name() string {
	return s.name
}

func (s *Scrapper) SetLanguage(lang string) {
	s.language = lang
}
