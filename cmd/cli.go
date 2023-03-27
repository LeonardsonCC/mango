package cmd

import (
	"fmt"
	"os"

	"github.com/LeonardsonCC/mango/internal/app/scrappers/mangalivre"
)

type Cli struct{}

func NewCli() *Cli {
	return &Cli{}
}

func (*Cli) Start() {
	args := os.Args[1:]
	if len(args) != 2 || args[0] == "" || args[1] == "" {
		fmt.Print(`usage: <binary> "anime-name" "10"`)
		return
	}

	name := args[0]
	chapter := args[1]

	s := mangalivre.NewScrapper()

	fmt.Println("searching manga")
	r, err := s.SearchManga(name)
	if err != nil {
		fmt.Printf("failed to find: %s", name)
		return
	}

	fmt.Println("searching chapter")
	c := s.SearchChapter(r[0].Url(), chapter)

	fmt.Println("downloading chapter")
	manga := s.Download(c[0].Url())

	filename := fmt.Sprintf("./%s.pdf", manga.Title)
	f, _ := os.Create(filename)
	defer f.Close()

	f.Write(manga.Buffer.Bytes())

}
