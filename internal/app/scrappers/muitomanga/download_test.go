package muitomanga_test

import (
	"testing"

	"github.com/LeonardsonCC/mango/internal/app/scrappers/muitomanga"
)

func TestDownloadMuitoManga(t *testing.T) {
	t.Run("download anime", func(t *testing.T) {
		type tc struct {
			url string
		}
		testCases := []tc{
			// {
			// 	url: "https://muitomanga.com/ler/kaguya-sama-wa-kokurasetai-doujin-ban/capitulo-27",
			// },
			{
				url: "https://muitomanga.com/ler/naruto/capitulo-700.11",
			},
		}

		s := muitomanga.NewScrapper()

		for _, c := range testCases {
			manga, _ := s.Download(c.url)

			t.Errorf("%+v", manga)
		}
	})
}
