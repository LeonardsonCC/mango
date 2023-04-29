package mangalivre

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/LeonardsonCC/mango/internal/app/scrappers"
)

func (s *Scrapper) SearchManga(query string) ([]*scrappers.SearchMangaResult, error) {
	var results []*scrappers.SearchMangaResult

	formData := url.Values{
		"search": {query},
	}

	client := &http.Client{}

	//Not working, the post data is not a form
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/lib/search/series.json", s.baseURL), strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-requested-with", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if _, ok := result["series"].([]interface{}); !ok {
		return results, nil
	}

	for _, r := range result["series"].([]interface{}) {
		re := r.(map[string]interface{})
		results = append(results, scrappers.NewSearchResult(
			re["label"].(string),
			re["cover"].(string),
			fmt.Sprintf("%s%s", s.baseURL, re["link"].(string)),
		))
	}

	return results, nil
}

func (s *Scrapper) SearchChapter(u, query string) ([]*scrappers.SearchChapterResult, error) {
	var chapters []*scrappers.SearchChapterResult

	re := regexp.MustCompile(`.*\/(\d*)`)
	match := re.FindStringSubmatch(u)
	mangaId := match[1]

	var i int
	for {
		i++

		formData := url.Values{
			"id_serie": {mangaId},
			"page":     {fmt.Sprintf("%d", i)},
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/series/chapters_list.json?%s", s.baseURL, formData.Encode()), nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("x-requested-with", "XMLHttpRequest")

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		result := make(map[string]interface{})
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Fatalln(err)
		}

		if _, ok := result["chapters"].(bool); ok {
			break
		}

		for _, chapter := range result["chapters"].([]interface{}) {
			link := ""
			c := chapter.(map[string]interface{})
			title := c["number"].(string)
			date := c["date"].(string)
			for _, r := range c["releases"].(map[string]interface{}) {
				l := r.(map[string]interface{})
				link = fmt.Sprintf("%s%s", s.baseURL, l["link"].(string))
			}
			ch := scrappers.NewSearchChapterResult(title, link, date)
			chapters = append(chapters, ch)
		}
	}

	var results []*scrappers.SearchChapterResult
	// searching using for (oldschool method)
	if query != "" {
		for _, chapter := range chapters {
			if strings.EqualFold(chapter.Title(), query) {
				results = append(results, chapter)
			}
		}
	} else {
		results = chapters
	}

	return results, nil
}
