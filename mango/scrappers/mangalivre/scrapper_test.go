package mangalivre_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/mango/scrappers/mangalivre"
)

func TestMangaLivre(t *testing.T) {
	t.Run("download range of mangas", func(t *testing.T) {
		s := mangalivre.NewScrapper()
		results, _ := s.SearchManga("tokyo ghoul")

		chapters, _ := s.SearchChapter(results[0].Url(), "")

		cc := make([]*scrappers.SearchChapterResult, len(chapters)+1)
		j := 0
		for i := len(chapters) - 1; i > -1; i-- {
			cc[j] = chapters[i]
			j++
		}

		wg := &sync.WaitGroup{}
		for _, c := range cc[0:20] {
			wg.Add(1)
			go func(c *scrappers.SearchChapterResult, wg *sync.WaitGroup) {
				manga, _ := s.Download(c.Url())

				filename := fmt.Sprintf("tokyo/%s.pdf", manga.Title)
				f, _ := os.Create(filename)
				f.ReadFrom(manga.Buffer)
				f.Close()
				fmt.Printf("Downloaded %s\n", filename)
				wg.Done()
			}(c, wg)
		}

		wg.Wait()
	})
}
