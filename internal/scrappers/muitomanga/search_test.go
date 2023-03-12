package muitomanga_test

import (
	"testing"

	"github.com/LeonardsonCC/mango/internal/scrappers/muitomanga"
)

func TestSearchMuitoManga(t *testing.T) {
	t.Run("search anime", func(t *testing.T) {
		type tc struct {
			q              string
			expectedUrl    string
			expectedImgUrl string
			expectedTitle  string
		}
		testCases := []tc{
			{
				q:              "boruto",
				expectedTitle:  "Boruto - Naruto Next Generations",
				expectedImgUrl: "https://imgs.muitomanga.com/posters/boruto-naruto-next-generations.jpg",
				expectedUrl:    "https://muitomanga.com/manga/boruto-naruto-next-generations",
			},
			{
				q:              "naruto",
				expectedTitle:  "Naruto",
				expectedImgUrl: "https://imgs.muitomanga.com/posters/naruto.jpg",
				expectedUrl:    "https://muitomanga.com/manga/naruto",
			},
		}

		s := muitomanga.NewScrapper()

		for _, c := range testCases {
			r := s.SearchAnime(c.q)

			if len(r) == 0 {
				t.Error("failed to get thumbnails")
			}

			worstAnime := r[0]
			if worstAnime == nil {
				t.Error("failed to get worst anime")
			}

			if worstAnime.Title() != c.expectedTitle {
				t.Errorf("found the wrong anime title. expected %s, but received %s", c.expectedTitle, worstAnime.Title())
			}

			if worstAnime.ImgUrl() != c.expectedImgUrl {
				t.Errorf("found the wrong anime img. expected %s, but received %s", c.expectedImgUrl, worstAnime.ImgUrl())
			}

			if worstAnime.Url() != c.expectedUrl {
				t.Errorf("found the wrong anime url. expected %s, but received %s", c.expectedUrl, worstAnime.Url())
			}
		}
	})

	t.Run("search chapter", func(t *testing.T) {
		s := muitomanga.NewScrapper()

		type tc struct {
			q                   string
			url                 string
			expectedTitle       string
			expectedUrl         string
			expectedAddedToSite string
		}

		testCases := []tc{
			{
				q:                   "",
				url:                 "https://muitomanga.com/manga/boruto-naruto-next-generations",
				expectedTitle:       "Capítulo #78 Boruto Explorer",
				expectedUrl:         "https://muitomanga.com/ler/boruto-naruto-next-generations/capitulo-78",
				expectedAddedToSite: "21/02/2023",
			},
			{
				q:                   "10",
				url:                 "https://muitomanga.com/manga/boruto-naruto-next-generations",
				expectedTitle:       "Capítulo #10 scansPROJECT",
				expectedUrl:         "https://muitomanga.com/ler/boruto-naruto-next-generations/capitulo-10",
				expectedAddedToSite: "20/09/2020",
			},
		}

		for _, c := range testCases {
			r := s.SearchChapter(c.url, c.q)

			if len(r) == 0 {
				t.Error("failed to get chapters")
			}

			firstChapter := r[0]

			if firstChapter.Title() != c.expectedTitle {
				t.Errorf("failed to get correct title. expected %s, received %s", c.expectedTitle, firstChapter.Title())
			}

			if firstChapter.Url() != c.expectedUrl {
				t.Errorf("failed to get correct Url. expected %s, received %s", c.expectedUrl, firstChapter.Url())
			}

			if firstChapter.AddedToSite() != c.expectedAddedToSite {
				t.Errorf("failed to get correct AddedToSite. expected %s, received %s", c.expectedAddedToSite, firstChapter.AddedToSite())
			}
		}
	})
}
