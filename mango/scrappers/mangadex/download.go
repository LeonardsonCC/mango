package mangadex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/LeonardsonCC/mango/internal/pkg/pdf"
	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/mango/scrappers/mangadex/entity"
	"github.com/LeonardsonCC/mango/pkg/mysync"
)

func (s *Scrapper) Download(chapter *scrappers.SearchChapterResult) (*scrappers.Manga, error) {
	resp, err := http.Get(fmt.Sprintf("%s/at-home/server/%s?forcePort443=false", s.apiBaseURL, chapter.ID()))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result := new(entity.ChapterPagesResult)
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	pagesToGet := make(map[int]string)
	for i, page := range result.Chapter.Data {
		pagesToGet[i] = fmt.Sprintf("%s/data/%s/%s", result.BaseURL, result.Chapter.Hash, page)
	}

	pages := s.collectPages(pagesToGet)

	w := bytes.NewBuffer([]byte{})

	err = pdf.GeneratePdf(pages, w)

	// print warnings from pdf generation
	if v, ok := err.(*pdf.Warnings); ok {
		if len(v.Warns) > 0 {
			fmt.Print(err.Error())
		}
	} else {
		return nil, err
	}

	m := scrappers.NewManga(pages, len(pages), chapter.Chapter()+" - "+chapter.Title(), w)

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
