package scrappers

type Scrapper interface {
	SearchManga(query string) ([]*SearchMangaResult, error)
	SearchChapter(url, query string) []*SearchChapterResult
	Download(url string) *Manga
}

type SearchMangaResult struct {
	title  string
	imgUrl string
	url    string
}

func NewSearchResult(title, imgUrl, url string) *SearchMangaResult {
	return &SearchMangaResult{
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

type SearchChapterResult struct {
	title       string
	url         string
	addedToSite string
}

func NewSearchChapterResult(title, url, addedToSite string) *SearchChapterResult {
	return &SearchChapterResult{
		title:       title,
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
