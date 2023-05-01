package muitomanga

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"sync"

	"github.com/LeonardsonCC/mango/internal/mango/scrappers"
	"github.com/LeonardsonCC/mango/internal/pkg/pdf"
	"github.com/LeonardsonCC/mango/pkg/mysync"
	"github.com/gocolly/colly"
)

func (s *Scrapper) Download(url string) (*scrappers.Manga, error) {
	pageNumber := s.getPageNumber(url)

	if pageNumber == 0 {
		return nil, errors.New("can't find the page number for this url")
	}

	name, chapter := s.extractProperties(url)

	cdn := s.findTheCdn(name, chapter)
	if cdn == "" {
		return nil, errors.New("can't find the cdn to download the pages for this url")
	}

	fullUrl := fmt.Sprintf("%s/%s/%s", cdn, name, chapter)

	pages := s.collectPages(fullUrl, pageNumber)

	w := bytes.NewBuffer([]byte{})

	err := pdf.GeneratePdf(pages, w)

	// print warnings from pdf generation
	if v, ok := err.(*pdf.Warnings); ok {
		if len(v.Warns) > 0 {
			fmt.Print(err.Error())
		}
	} else {
		return nil, err
	}

	m := scrappers.NewManga(pages, pageNumber, fmt.Sprintf("%s_%s", name, chapter), w)

	return m, nil
}

// getPageNumber goes to manga page, and gets the page number
func (s *Scrapper) getPageNumber(url string) int {
	var pageNumber int
	s.Colly.OnHTML(".select_paged", func(e *colly.HTMLElement) {
		found := e.ChildText("option")
		re := regexp.MustCompile(`\/ (\d.)`)
		match := re.FindStringSubmatch(found)

		n, err := strconv.Atoi(match[1])
		if err != nil {
			return
		}

		pageNumber = n
	})

	s.Colly.Visit(url)

	return pageNumber
}

// extractProperties extract the title and chapter from url
func (s *Scrapper) extractProperties(url string) (string, string) {
	re := regexp.MustCompile(`\/ler\/(.+)\/capitulo-(\d+(?:\.\d+)?)`)
	match := re.FindStringSubmatch(url)

	return match[1], match[2]
}

// TODO: improve this function performance
// maybe using channels
// collectPages goes through each page of the manga and download it
// reutrning the map[page]image
func (s *Scrapper) collectPages(url string, pageNumber int) map[int][]byte {
	pages := mysync.NewMap(map[int][]byte{})

	wg := &sync.WaitGroup{}

	for i := 0; i <= pageNumber; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			resp, err := http.Get(fmt.Sprintf("%s/%d.jpg", url, i))
			if err != nil {
				return
			}

			if resp.StatusCode == http.StatusNotFound {
				return
			}

			page, err := io.ReadAll(resp.Body)
			if err != nil {
				return
			}

			pages.Store(i, page)
		}(i)
	}
	wg.Wait()

	return pages.Map()
}

// findTheCdn tries to find the chapter in the known CDNs
// and returns the images URL
func (s *Scrapper) findTheCdn(name, chapter string) string {
	for _, url := range s.imagesBaseURLs {
		u := fmt.Sprintf("%s/imgs", url)

		resp, err := http.Get(fmt.Sprintf("%s/%s/%s/1.jpg", u, name, chapter))
		if err != nil {
			continue
		}

		if resp.StatusCode == http.StatusOK {
			return u
		}
	}

	return ""
}
