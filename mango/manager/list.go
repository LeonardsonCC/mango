package manager

import (
	"sync"

	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/pkg/mysync"
)

func (m *Manager) ListChapters(name string) (map[string][]*scrappers.SearchChapterResult, map[string]error) {
	results := mysync.NewMap(
		make(map[string][]*scrappers.SearchChapterResult, len(m.scrappers)),
	)
	errs := mysync.NewMap(
		make(map[string]error, len(m.scrappers)),
	)

	var wg sync.WaitGroup

	for k, s := range m.scrappers {
		k := k
		s := s

		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := s.SearchManga(name)
			if err != nil {
				errs.Store(k, err)
				results.Store(k, make([]*scrappers.SearchChapterResult, 0))
				return
			}

			if len(r) == 0 {
				results.Store(k, make([]*scrappers.SearchChapterResult, 0))
				return
			}

			manga := r[0]
			chapters, err := s.SearchChapter(manga, "")
			if err != nil {
				errs.Store(k, err)
				results.Store(k, make([]*scrappers.SearchChapterResult, 0))
				return
			}

			results.Store(k, chapters)
		}()
	}

	wg.Wait()

	return results.Map(), errs.Map()
}
