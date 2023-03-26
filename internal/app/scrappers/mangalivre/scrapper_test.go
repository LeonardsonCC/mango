package mangalivre_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/LeonardsonCC/mango/internal/app/scrappers/mangalivre"
)

func TestMangaLivre(t *testing.T) {
	t.Run("download range of mangas", func(t *testing.T) {
		s := mangalivre.NewScrapper()
		results := s.SearchManga("relife")

		chapters := s.SearchChapter(results[0].Url(), "")

		for _, c := range chapters {
			manga := s.Download(c.Url())

			filename := fmt.Sprintf("relife/%s.pdf", manga.Title)
			f, _ := os.Create(filename)
			f.ReadFrom(manga.Buffer)
			f.Close()
		}
	})
}
