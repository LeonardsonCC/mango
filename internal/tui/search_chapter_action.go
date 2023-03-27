package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type chapterSearchDone []list.Item

func (m *mangoTui) searchChapterAction() tea.Cmd {
	return func() tea.Msg {
		// should return error to handle
		results, err := m.scrapper.SearchChapter(m.mangaUrl, m.chapter)
		if err != nil {
			// TODO: handle error
			return chapterSearchDone([]list.Item{})
		}

		items := make([]list.Item, len(results))
		for i, r := range results {
			items[i] = item{text: r.Title(), value: r.Url()}
		}

		return chapterSearchDone(items)
	}
}
