package manager

import (
	"github.com/LeonardsonCC/mango/internal/mango/scrappers"
	"github.com/LeonardsonCC/mango/pkg/mysync"
	"golang.org/x/sync/errgroup"
)

func (m *Manager) ListChapters(name string) (map[string][]*scrappers.SearchChapterResult, error) {
	results := mysync.NewMap(
		make(map[string][]*scrappers.SearchChapterResult, len(m.scrappers)),
	)

	g := new(errgroup.Group)

	for k, s := range m.scrappers {
		k := k
		s := s
		g.Go(func() error {
			r, err := s.SearchManga(name)
			if err != nil {
				return err
			}

			if len(r) == 0 {
				return nil
			}

			manga := r[0]
			chapters, err := s.SearchChapter(manga.Url(), "")
			if err != nil {
				return err
			}

			results.Store(k, chapters)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results.Map(), nil
}
