package mangadex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/mango/scrappers/mangadex/entity"
)

func (s *Scrapper) SearchManga(query string) ([]*scrappers.SearchMangaResult, error) {
	var results []*scrappers.SearchMangaResult

	formData := url.Values{
		"title":            {query},
		"limit":            {"5"},
		"contentRating[]":  {"safe", "suggestive", "erotica"},
		"includes[]":       {"cover_art"},
		"order[relevance]": {"desc"},
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/manga?%s", s.apiBaseURL, formData.Encode()), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := new(entity.MangaDexSearchResult)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	for _, manga := range result.Data {
		var coverFileName string
		for _, relationship := range manga.Relationships {
			if relationship.Type == "cover_art" {
				coverFileName = relationship.Attributes.FileName
				break
			}
		}

		results = append(results,
			scrappers.NewSearchResult(
				manga.ID,
				manga.Attributes.Title.EN,
				fmt.Sprintf("%s/covers/%s/%s", s.baseURL, manga.ID, coverFileName),
				fmt.Sprintf("%s/title/%s", s.baseURL, manga.ID),
			))
	}

	return results, nil
}

func (s *Scrapper) SearchChapter(manga *scrappers.SearchMangaResult, query string) ([]*scrappers.SearchChapterResult, error) {
	var results []*scrappers.SearchChapterResult

	limit := 100
	i := 0

outer:
	for {
		formData := url.Values{
			"limit":                {fmt.Sprintf("%d", limit)},
			"includes[]":           {"scanlation_group", "user"},
			"order[volume]":        {"desc"},
			"order[chapter]":       {"asc"},
			"offset":               {fmt.Sprintf("%d", i*limit)},
			"contentRating[]":      {"safe", "suggestive", "erotica"},
			"translatedLanguage[]": {s.language},
		}

		client := &http.Client{}

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/manga/%s/feed?%s", s.apiBaseURL, manga.ID(), formData.Encode()), nil)
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		result := new(entity.MangaDexSearchChapterResult)
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return nil, err
		}

		for _, chapter := range result.Data {
			chapterResult := scrappers.NewSearchChapterResult(
				chapter.ID,
				chapter.Attributes.Title,
				chapter.Attributes.Chapter,
				fmt.Sprintf("%s/chapter/%s", s.baseURL, chapter.ID),
				chapter.Attributes.CreatedAt,
			)

			if query == chapter.Attributes.Chapter {
				results = []*scrappers.SearchChapterResult{
					chapterResult,
				}
				break outer
			}

			results = append(results, chapterResult)
		}

		if i*limit >= result.Total {
			break
		}
		i++
	}

	return results, nil
}
