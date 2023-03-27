package muitomanga

import (
	"fmt"
	"strings"

	"github.com/LeonardsonCC/mango/internal/app/scrappers"
	"github.com/gocolly/colly"
)

// Search receives the string query and will search on muitomanga.com
// the manga with that name
func (s *Scrapper) SearchManga(query string) ([]*scrappers.SearchMangaResult, error) {
	var results []*scrappers.SearchMangaResult

	s.Colly.OnHTML(".anime", func(e *colly.HTMLElement) {
		img := e.ChildAttr("img", "src")
		link := e.ChildText("h3")
		url := e.ChildAttr("a", "href")

		r := scrappers.NewSearchResult(link, img, fmt.Sprintf("%s%s", s.baseURL, url))

		results = append(results, r)
	})

	err := s.Colly.Visit(fmt.Sprintf("%s/buscar?q=%s", s.baseURL, query))
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *Scrapper) SearchChapter(url, query string) ([]*scrappers.SearchChapterResult, error) {
	var chapters []*scrappers.SearchChapterResult

	s.Colly.OnHTML(".single-chapter", func(e *colly.HTMLElement) {
		title := e.ChildText("a")
		url := e.ChildAttr("a", "href")
		addedToSite := e.ChildText("small")

		r := scrappers.NewSearchChapterResult(title, fmt.Sprintf("%s%s", s.baseURL, url), addedToSite)

		chapters = append(chapters, r)
	})

	err := s.Colly.Visit(url)
	if err != nil {
		return nil, err
	}

	if query == "" {
		return chapters, nil
	}

	var results []*scrappers.SearchChapterResult

	// searching using for (oldschool method)
	for _, chapter := range chapters {
		if strings.Contains(chapter.Title(), fmt.Sprintf("#%s", query)) {
			results = append(results, chapter)
		}
	}

	return results, nil
}
