package tui

import tea "github.com/charmbracelet/bubbletea"

type loading bool

func setLoading(newValue bool) tea.Cmd {
	return func() tea.Msg {
		return loading(newValue)
	}
}
