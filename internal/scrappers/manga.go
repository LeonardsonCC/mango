package scrappers

import "io"

type Manga struct {
	PagesQty int
	Pages    []*Page
	Title    string
	Writer   io.Writer
}

func NewManga(pages map[int][]byte, pagesQty int, title string, w io.Writer) *Manga {
	p := make([]*Page, len(pages))
	for _, page := range pages {
		p = append(p, NewPage(page))
	}

	return &Manga{
		Pages:    p,
		Title:    title,
		PagesQty: pagesQty,
		Writer:   w,
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
