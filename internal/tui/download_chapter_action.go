package tui

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *mangoTui) downloadChapterAction() tea.Cmd {
	return func() tea.Msg {
		manga, err := m.scrapper.Download(m.chapterUrl)
		if err != nil {
			// TODO: handle error better
			return loading(false)
		}

		filename := fmt.Sprintf("./%s.pdf", manga.Title)
		f, _ := os.Create(filename)
		f.ReadFrom(manga.Buffer)
		f.Close()

		c := exec.Command("xdg-open", filename)
		c.Run()

		return loading(false)
	}
}
