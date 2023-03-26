package cmd

import (
	"fmt"
	"os"

	"github.com/LeonardsonCC/mango/internal/app/scrappers/mangalivre"
	"github.com/LeonardsonCC/mango/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

type Tui struct{}

func NewTui() *Tui {
	return &Tui{}
}

func (*Tui) Start() {
	s := mangalivre.NewScrapper()

	p := tea.NewProgram(tui.InitTui(s), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
