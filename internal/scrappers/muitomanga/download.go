package muitomanga

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"

	"github.com/LeonardsonCC/mango/internal/pdf"
	"github.com/LeonardsonCC/mango/internal/scrappers"
	"github.com/gocolly/colly"
)

type listOfPages struct {
	sync.Mutex
	m map[int][]byte
}

func (l *listOfPages) Store(key int, value []byte) {
	l.Lock()
	defer l.Unlock()

	l.m[key] = value
}

func (l *listOfPages) Get(key int) []byte {
	l.Lock()
	defer l.Unlock()

	if v, has := l.m[key]; has {
		return v
	}
	return nil
}

func (s *Scrapper) Download(url string) *scrappers.Manga {
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

	if pageNumber == 0 {
		return nil
	}

	re := regexp.MustCompile(`\/ler\/(.+)\/capitulo-(\d+(?:\.\d+)?)`)
	match := re.FindStringSubmatch(url)

	name := match[1]
	chapter := match[2]

	cdn := s.findTheCdn(name, chapter)
	if cdn == "" {
		return nil
	}

	fullUrl := fmt.Sprintf("%s/%s/%s", cdn, name, chapter)

	pages := s.collectPages(fullUrl, pageNumber)

	filename := fmt.Sprintf("./%s_%s.pdf", name, chapter)
	f, _ := os.Create(filename)
	defer f.Close()

	pdf.GeneratePdf(pages, f)

	m := scrappers.NewManga(pages, pageNumber, fmt.Sprintf("%s_%s", name, chapter), f)

	return m
}

// TODO: improve this function performance
// maybe using channels
func (s *Scrapper) collectPages(url string, pageNumber int) map[int][]byte {
	pages := &listOfPages{
		m: map[int][]byte{},
	}

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

	return pages.m
}

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
