package mangalivre_test

import (
	"testing"

	"github.com/LeonardsonCC/mango/mango/scrappers/mangalivre"
)

func TestDownloadmangalivre(t *testing.T) {
	t.Run("download anime", func(t *testing.T) {
		type tc struct {
			url string
		}
		testCases := []tc{
			{
				url: "https://mangalivre.net/ler/null/online/92635/null",
			},
		}

		s := mangalivre.NewScrapper()

		for _, c := range testCases {
			manga, _ := s.Download(c.url)

			t.Errorf("%+v", manga)
		}
	})
}
