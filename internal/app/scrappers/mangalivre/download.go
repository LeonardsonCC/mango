package mangalivre

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"

	"github.com/LeonardsonCC/mango/internal/app/scrappers"
	"github.com/LeonardsonCC/mango/internal/pkg/pdf"
	"github.com/LeonardsonCC/mango/pkg/mysync"
)

func (s *Scrapper) Download(u string) (*scrappers.Manga, error) {
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	readerToken := getReaderToken(res)
	releaseId := getReleaseId(res)

	formData := url.Values{
		"key": {readerToken},
	}

	client := &http.Client{}

	//Not working, the post data is not a form
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/leitor/pages/%s.json", s.baseURL, releaseId), strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	p := make(map[int]string)
	for i, result := range result["images"].([]interface{}) {
		r := result.(map[string]interface{})
		p[i] = r["legacy"].(string)
	}

	images := s.collectPages(p)

	w := bytes.NewBuffer([]byte{})

	err = pdf.GeneratePdf(images, w)

	// print warnings from pdf generation
	if v, ok := err.(*pdf.Warnings); ok {
		if len(v.Warns) > 0 {
			fmt.Print(err.Error())
		}
	} else {
		return nil, err
	}

	m := scrappers.NewManga(images, len(images), getMangaName(res), w)

	return m, nil
}

func (s *Scrapper) collectPages(p map[int]string) map[int][]byte {
	pages := mysync.NewMap(map[int][]byte{})

	wg := &sync.WaitGroup{}

	for i, page := range p {
		wg.Add(1)
		go func(wg *sync.WaitGroup, page string, i int) {
			resp, err := http.Get(page)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			res, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			pages.Store(i, res)
			wg.Done()
		}(wg, page, i)
	}
	wg.Wait()

	return pages.Map()
}

func getReaderToken(body []byte) string {
	re := regexp.MustCompile(`window\.READER_TOKEN = '(.*)';`)
	match := re.FindStringSubmatch(string(body))
	return match[1]
}

func getReleaseId(body []byte) string {
	re := regexp.MustCompile(`window\.READER_ID_RELEASE = '(.*)';`)
	match := re.FindStringSubmatch(string(body))
	return match[1]
}

func getMangaName(body []byte) string {
	re := regexp.MustCompile(`<title>(.*)<\/title>`)
	match := re.FindStringSubmatch(string(body))
	return match[1]
}
