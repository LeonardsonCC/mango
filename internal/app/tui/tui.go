package tui

import (
	"fmt"
	"io"

	"github.com/LeonardsonCC/mango/internal/app/scrappers"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	itemStyle         = lipgloss.NewStyle()
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

type item struct {
	text  string
	value string
}

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := i.text

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render(fmt.Sprintf("> %s", s))
		}
	}

	fmt.Fprint(w, fn(str))
}

type Step int

const (
	StepSearchManga Step = iota
	StepListManga
	StepSearchChapter
	StepListChapter
	StepDownloadingChapter
)

type mangoTui struct {
	textInput  textinput.Model
	list       list.Model
	step       Step
	manga      string
	chapter    string
	mangaUrl   string
	chapterUrl string
	loading    bool

	scrapper scrappers.Scrapper
}

func InitTui(scrapper scrappers.Scrapper) mangoTui {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 60

	width, height, _ := term.GetSize(0)

	l := list.New([]list.Item{}, itemDelegate{}, width, height)
	// l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return mangoTui{
		textInput: ti,
		list:      l,
		step:      StepSearchManga,
		scrapper:  scrapper,
	}
}

func (m mangoTui) Init() tea.Cmd {
	return textinput.Blink
}

func (m mangoTui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			switch m.step {
			case StepSearchManga, StepSearchChapter:
				if m.step == StepSearchManga {
					m.manga = m.textInput.Value()
					m.step = StepListManga
					m.loading = true
					return m, m.searchMangaAction()
				} else if m.step == StepSearchChapter {
					m.chapter = m.textInput.Value()
					m.step = StepListChapter
					m.loading = true
					return m, m.searchChapterAction()
				}
				m.textInput.SetValue("")

			case StepListManga, StepListChapter:
				selected := m.list.SelectedItem()
				i, _ := selected.(item)

				if m.step == StepListManga {
					m.mangaUrl = i.value
					m.step = StepSearchChapter
				} else if m.step == StepListChapter {
					m.chapterUrl = i.value
					m.loading = true
					return m, m.downloadChapterAction()
				}
			}

		case tea.KeyCtrlC, tea.KeyEsc:
			switch m.step {
			case StepListManga:
				m.step = StepSearchManga
				m.textInput.SetValue("")
			case StepSearchChapter:
				m.step = StepListManga
				m.textInput.SetValue("")
			case StepListChapter:
				m.step = StepSearchChapter
			default:
				return m, tea.Quit
			}
		}
	case mangaSearchDone:
		m.list.SetItems(msg)
		m.textInput.SetValue("")
		return m, setLoading(false)
	case chapterSearchDone:
		m.list.SetItems(msg)
		return m, setLoading(false)
	case loading:
		m.loading = bool(msg)
		return m, cmd
	}

	if m.step == StepListManga || m.step == StepListChapter {
		m.list, _ = m.list.Update(msg)
	}

	if m.step == StepSearchManga || m.step == StepSearchChapter {
		m.textInput, _ = m.textInput.Update(msg)
	}
	return m, cmd
}

func (m mangoTui) View() string {
	var s string

	if m.loading {
		s = "Loading...\n"
		return s
	}

	switch m.step {
	case StepSearchManga:
		s += "Search a manga\n"
		s += m.textInput.View()

	case StepListManga:
		s += m.list.View()

	case StepSearchChapter:
		s += "What chapter do you want?\n"
		s += m.textInput.View()

	case StepListChapter:
		s += m.list.View()
	}

	return s
}
