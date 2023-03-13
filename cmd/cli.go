package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/LeonardsonCC/mango/internal/app/scrappers/muitomanga"
)

func Start() {
	args := os.Args[1:]
	if len(args) != 2 || args[0] == "" || args[1] == "" {
		fmt.Print(`usage: <binary> "anime-name" "10"`)
		return
	}

	name := args[0]
	chapter := args[1]

	s := muitomanga.NewScrapper()

	log.Print("searching manga")
	r := s.SearchManga(name)

	log.Print("searching chapter")
	c := s.SearchChapter(r[0].Url(), chapter)

	log.Print("downloading chapter")
	manga := s.Download(c[0].Url())

	filename := fmt.Sprintf("./%s.pdf", manga.Title)
	f, _ := os.Create(filename)
	defer f.Close()

	f.Write(manga.Buffer.Bytes())

}
