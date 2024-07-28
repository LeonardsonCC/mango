package manager

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/LeonardsonCC/mango/mango/scrappers"
	"github.com/LeonardsonCC/mango/pkg/mysync"
	"golang.org/x/sync/errgroup"
)

func (m *Manager) DownloadChapter(name, chapter string) error {
	var manga *scrappers.Manga

	results, _ := m.Search(name)

	chapters := mysync.NewMap(
		make(map[string][]*scrappers.SearchChapterResult, len(m.scrappers)),
	)
	for k, s := range results {
		if len(s) > 0 {
			scrapper := m.scrappers[k]
			manga := s[0]
			found, err := scrapper.SearchChapter(manga, chapter)
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
			m, err := m.scrappers[k].Download(scr[0])
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

	filename := path.Join(m.output, fmt.Sprintf("%s.pdf", MakeFileName(manga.Title)))
	f, _ := os.Create(filename)
	defer f.Close()

	_, err := f.Write(manga.Buffer.Bytes())
	if err != nil {
		return fmt.Errorf("failed to save pdf")
	}

	return nil
}

func (m *Manager) DownloadAllChapters(mangaName string) error {
	results := mysync.NewMap(
		make(map[string][]*scrappers.SearchMangaResult, len(m.scrappers)),
	)

	g := new(errgroup.Group)

	for k, s := range m.scrappers {
		k := k
		s := s

		s.SetInfoChannel(m.infoChannel)

		g.Go(func() error {
			r, err := s.SearchManga(mangaName)
			if err != nil {
				return err
			}

			results.Store(k, r)
			m.infoChannel <- fmt.Sprintf("manga found: %s", r[0].Title())
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
			found, err := scrapper.SearchChapter(manga, "")
			if err != nil {
				return err
			}

			m.infoChannel <- fmt.Sprintf("chapters found: %d", len(found))
			chapters.Store(k, found)
		}
	}

	c := chapters.Map()
	if len(c) == 0 {
		return fmt.Errorf("none of the scrappers found the manga %s", mangaName)
	}

	g = new(errgroup.Group)
	g.SetLimit(40)

	for k, scr := range c {
		if len(scr) > 0 {
			for _, chap := range scr {
				chap := chap
				g.Go(func() error {
					mang, err := m.scrappers[k].Download(chap)
					if err != nil {
						return err
					}
					m.infoChannel <- fmt.Sprintf("downloaded: %s", chap.Chapter())

					filename := path.Join(m.output, fmt.Sprintf("%s.pdf", MakeFileName(mang.Title)))
					f, err := os.Create(filename)
					if err != nil {
						m.infoChannel <- fmt.Sprintf("failed to create file: %s", filename)
					}

					defer f.Close()

					n, err := f.Write(mang.Buffer.Bytes())
					if err != nil || n == 0 {
						m.infoChannel <- fmt.Sprintf("failed to write pdf: %s", err)
						return fmt.Errorf("failed to save pdf")
					}
					if n == 431 {
						m.infoChannel <- mang.Buffer.String()
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

var invalidNameRegexp = regexp.MustCompile(`[/\\?%*:|"<>]`)

func MakeFileName(title string) string {
	return invalidNameRegexp.ReplaceAllString(title, "")
}
