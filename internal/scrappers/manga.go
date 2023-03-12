package scrappers

import (
	"bytes"
)

type Manga struct {
	PagesQty int
	Pages    []*Page
	Title    string
	Buffer   *bytes.Buffer
}

func NewManga(pages map[int][]byte, pagesQty int, title string, w *bytes.Buffer) *Manga {
	p := make([]*Page, len(pages))
	for _, page := range pages {
		p = append(p, NewPage(page))
	}

	return &Manga{
		Pages:    p,
		Title:    title,
		PagesQty: pagesQty,
		Buffer:   w,
	}
}

type Page struct {
	Content []byte
}

func NewPage(p []byte) *Page {
	return &Page{
		Content: p,
	}
}
