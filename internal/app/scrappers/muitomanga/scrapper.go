package muitomanga

import (
	"fmt"

	"github.com/LeonardsonCC/mango/internal/app/scrappers"
	"github.com/gocolly/colly"
)

const ScrapperName = "MuitoManga"

var mainURL = "muitomanga.com"
var imagesURL = []string{
	"imgs.muitomanga.com",
	"imgs2.muitomanga.com",
}

type Scrapper struct {
	name           string
	baseURL        string
	imagesBaseURLs []string
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
		imagesBaseURLs: imgsBaseUrls,
	}
}

func (s *Scrapper) Name() string {
	return s.name
}
