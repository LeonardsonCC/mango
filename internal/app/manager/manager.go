package manager

import (
	"github.com/LeonardsonCC/mango/internal/app/scrappers"
	"github.com/LeonardsonCC/mango/internal/app/scrappers/mangalivre"
	"github.com/LeonardsonCC/mango/internal/app/scrappers/muitomanga"
	"github.com/LeonardsonCC/mango/pkg/mysync"
	"golang.org/x/sync/errgroup"
)

type Manager struct {
	scrappers []scrappers.Scrapper
	state     *ManagerState
}

type ManagerState struct {
	scrapper scrappers.Scrapper
	manga    *scrappers.SearchMangaResult
	chapter  *scrappers.SearchChapterResult
}

func NewManager() *Manager {
	return &Manager{
		scrappers: []scrappers.Scrapper{
			mangalivre.NewScrapper(),
			muitomanga.NewScrapper(),
		},
		state: &ManagerState{},
	}
}

func (m *Manager) SearchManga(query string) ([]*scrappers.SearchMangaResult, error) {
	w := new(errgroup.Group)

	results := mysync.NewSlice(make([]*scrappers.SearchMangaResult, 0, 30))

	for _, s := range m.scrappers {
		// you know, a workaround for the goroutine
		s := s
		w.Go(func() error {
			r, err := s.SearchManga(query)

			for _, v := range r {
				results.Append(v)
			}

			return err
		})
	}

	if err := w.Wait(); err != nil {
		return nil, err
	}

	return results.Slice(), nil
}

func (m *Manager) SearchChapter(query string) ([]*scrappers.SearchChapterResult, error) {
	return nil, nil
}

func (m *ManagerState) clear() {
	m.scrapper = nil
	m.manga = nil
	m.chapter = nil
}
