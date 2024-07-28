package scrappers

type Scrapper interface {
	SearchManga(query string) ([]*SearchMangaResult, error)
	SearchChapter(manga *SearchMangaResult, query string) ([]*SearchChapterResult, error)
	Download(chapter *SearchChapterResult) (*Manga, error)
	Name() string
	SetLanguage(language string)
	SetInfoChannel(info chan string)
}

type SearchMangaResult struct {
	title  string
	imgUrl string
	url    string
	id     string
}

func NewSearchResult(id, title, imgUrl, url string) *SearchMangaResult {
	return &SearchMangaResult{
		id:     id,
		title:  title,
		imgUrl: imgUrl,
		url:    url,
	}
}

func (s *SearchMangaResult) Title() string {
	return s.title
}

func (s *SearchMangaResult) ImgUrl() string {
	return s.imgUrl
}

func (s *SearchMangaResult) Url() string {
	return s.url
}

func (s *SearchMangaResult) ID() string {
	return s.id
}

type SearchChapterResult struct {
	id          string
	title       string
	chapter     string
	url         string
	addedToSite string
}

func NewSearchChapterResult(id, title, chapter, url, addedToSite string) *SearchChapterResult {
	return &SearchChapterResult{
		id:          id,
		title:       title,
		chapter:     chapter,
		url:         url,
		addedToSite: addedToSite,
	}
}

func (s *SearchChapterResult) Title() string {
	return s.title
}

func (s *SearchChapterResult) Url() string {
	return s.url
}

func (s *SearchChapterResult) AddedToSite() string {
	return s.addedToSite
}

func (s *SearchChapterResult) ID() string {
	return s.id
}

func (s *SearchChapterResult) Chapter() string {
	return s.chapter
}
