package scrappers

type Scrapper interface {
	SearchAnime(query string) []*SearchAnimeResult
	SearchChapter(url, query string) []*SearchChapterResult
	Download(url string) *Manga
}

type SearchAnimeResult struct {
	title  string
	imgUrl string
	url    string
}

func NewSearchResult(title, imgUrl, url string) *SearchAnimeResult {
	return &SearchAnimeResult{
		title:  title,
		imgUrl: imgUrl,
		url:    url,
	}
}

func (s *SearchAnimeResult) Title() string {
	return s.title
}

func (s *SearchAnimeResult) ImgUrl() string {
	return s.imgUrl
}

func (s *SearchAnimeResult) Url() string {
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
