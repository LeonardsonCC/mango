package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type mangaSearchDone []list.Item

func (m *mangoTui) searchMangaAction() tea.Cmd {
	return func() tea.Msg {
		// should return error to handle
		results := m.scrapper.SearchManga(m.manga)

		items := make([]list.Item, len(results))
		for i, r := range results {
			items[i] = item{text: r.Title(), value: r.Url()}
		}

		return mangaSearchDone(items)
	}
}
