package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LeonardsonCC/mango/internal/scrappers/muitomanga"
)

func main() {
	s := muitomanga.NewScrapper()

	log.Print("searching anime")
	r := s.SearchAnime("boku no hero")

	log.Print("searching chapter")
	c := s.SearchChapter(r[0].Url(), "10")

	log.Print("downloading chapter")
	manga := s.Download(c[0].Url())

	filename := fmt.Sprintf("./%s.pdf", manga.Title)
	f, _ := os.Create(filename)
	defer f.Close()

	f.Write(manga.Buffer.Bytes())
}
