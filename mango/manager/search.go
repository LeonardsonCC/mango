package manager

import (
	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/pkg/mysync"
	"golang.org/x/sync/errgroup"
)

func (m *Manager) Search(name string) (map[string][]*scrappers.SearchMangaResult, error) {
	results := mysync.NewMap(
		make(map[string][]*scrappers.SearchMangaResult, len(m.scrappers)),
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

			results.Store(k, r)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results.Map(), nil
}
