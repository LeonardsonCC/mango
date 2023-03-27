package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type mangaSearchDone []list.Item

func (m *mangoTui) searchMangaAction() tea.Cmd {
	return func() tea.Msg {
		// should return error to handle
		results, err := m.scrapper.SearchManga(m.manga)
		if err != nil {
			// TODO: handle error
			return mangaSearchDone([]list.Item{})
		}

		items := make([]list.Item, len(results))
		for i, r := range results {
			items[i] = item{text: r.Title(), value: r.Url()}
		}

		return mangaSearchDone(items)
	}
}
