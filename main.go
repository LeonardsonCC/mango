package main

import (
	"log"

	"github.com/LeonardsonCC/mango/scrappers/muitomanga"
)

func main() {
	s := muitomanga.NewScrapper()

	log.Print("searching anime")
	r := s.SearchAnime("boku no hero")

	log.Print("searching chapter")
	c := s.SearchChapter(r[0].Url(), "10")

	log.Print("downloading chapter")
	s.Download(c[0].Url())
}
