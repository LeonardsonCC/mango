package manager

import (
	"sync"

	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/pkg/mysync"
)

func (m *Manager) Search(name string) (map[string][]*scrappers.SearchMangaResult, map[string]error) {
	results := mysync.NewMap(
		make(map[string][]*scrappers.SearchMangaResult, len(m.scrappers)),
	)
	errs := mysync.NewMap(
		make(map[string]error, len(m.scrappers)),
	)

	var wg sync.WaitGroup

	for k, s := range m.scrappers {
		k, s := k, s

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			r, err := s.SearchManga(name)
			if err != nil {
				errs.Store(k, err)
				results.Store(k, make([]*scrappers.SearchMangaResult, 0))
				return
			}

			results.Store(k, r)
		}(&wg)
	}

	wg.Wait()

	return results.Map(), errs.Map()
}
