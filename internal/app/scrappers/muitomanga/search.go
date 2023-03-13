package muitomanga

import (
	"fmt"
	"strings"

	"github.com/LeonardsonCC/mango/internal/app/scrappers"
	"github.com/gocolly/colly"
)

// Search receives the string query and will search on muitomanga.com
// the manga with that name
func (s *Scrapper) SearchAnime(query string) []*scrappers.SearchAnimeResult {
	var results []*scrappers.SearchAnimeResult

	s.Colly.OnHTML(".anime", func(e *colly.HTMLElement) {
		img := e.ChildAttr("img", "src")
		link := e.ChildText("h3")
		url := e.ChildAttr("a", "href")

		r := scrappers.NewSearchResult(link, img, fmt.Sprintf("%s%s", s.baseURL, url))

		results = append(results, r)
	})

	s.Colly.Visit(fmt.Sprintf("%s/buscar?q=%s", s.baseURL, query))

	return results
}

func (s *Scrapper) SearchChapter(url, query string) []*scrappers.SearchChapterResult {
	var chapters []*scrappers.SearchChapterResult

	s.Colly.OnHTML(".single-chapter", func(e *colly.HTMLElement) {
		title := e.ChildText("a")
		url := e.ChildAttr("a", "href")
		addedToSite := e.ChildText("small")

		r := scrappers.NewSearchChapterResult(title, fmt.Sprintf("%s%s", s.baseURL, url), addedToSite)

		chapters = append(chapters, r)
	})

	s.Colly.Visit(url)

	if query == "" {
		return chapters
	}

	var results []*scrappers.SearchChapterResult

	// searching using for (oldschool method)
	for _, chapter := range chapters {
		if strings.Contains(chapter.Title(), fmt.Sprintf("#%s", query)) {
			results = append(results, chapter)
		}
	}

	return results
}
