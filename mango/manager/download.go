package manager

import (
	"fmt"
	"os"
	"path"

	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/pkg/mysync"
	"golang.org/x/sync/errgroup"
)

func (m *Manager) DownloadChapter(name, chapter string) error {
	var manga *scrappers.Manga

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
		return err
	}

	chapters := mysync.NewMap(
		make(map[string][]*scrappers.SearchChapterResult, len(m.scrappers)),
	)
	for k, s := range results.Map() {
		if len(s) > 0 {
			scrapper := m.scrappers[k]
			manga := s[0]
			found, err := scrapper.SearchChapter(manga.Url(), chapter)
			if err != nil {
				return err
			}

			chapters.Store(k, found)
		}
	}

	c := chapters.Map()
	if len(c) == 0 {
		return fmt.Errorf("none of the scrappers found the manga %s", name)
	}

	for k, scr := range c {
		if len(scr) > 0 {
			m, err := m.scrappers[k].Download(scr[0].Url())
			if err != nil {
				continue
			}
			manga = m
			break
		}
	}

	if manga == nil {
		return fmt.Errorf("failed to find chapter %s in manga %s", name, chapter)
	}

	filename := path.Join(m.output, fmt.Sprintf("%s.pdf", manga.Title))
	f, _ := os.Create(filename)
	defer f.Close()

	_, err := f.Write(manga.Buffer.Bytes())
	if err != nil {
		return fmt.Errorf("failed to save pdf")
	}

	return nil
}

func (m *Manager) DownloadManga(name string) error {
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
		return err
	}

	chapters := mysync.NewMap(
		make(map[string][]*scrappers.SearchChapterResult, len(m.scrappers)),
	)
	for k, s := range results.Map() {
		if len(s) > 0 {
			scrapper := m.scrappers[k]
			manga := s[0]
			// get all chapters
			found, err := scrapper.SearchChapter(manga.Url(), "")
			if err != nil {
				return err
			}

			chapters.Store(k, found)
		}
	}

	c := chapters.Map()
	if len(c) == 0 {
		return fmt.Errorf("none of the scrappers found the manga %s", name)
	}

	g = new(errgroup.Group)
	g.SetLimit(20)

	for k, scr := range c {
		if len(scr) > 0 {
			for _, chap := range scr {
				chap := chap
				g.Go(func() error {
					mang, err := m.scrappers[k].Download(chap.Url())
					if err != nil {
						return err
					}

					filename := path.Join(m.output, fmt.Sprintf("%s.pdf", mang.Title))
					f, _ := os.Create(filename)
					defer f.Close()

					_, err = f.Write(mang.Buffer.Bytes())
					if err != nil {
						return fmt.Errorf("failed to save pdf")
					}
					return nil
				})
			}
			break
		}
	}

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
