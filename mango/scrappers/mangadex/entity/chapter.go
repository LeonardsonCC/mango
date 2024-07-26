package entity

type MangaDexSearchChapterResult struct {
	Data     []MangaDexSearchChapterResultData `json:"data,omitempty"`
	Limit    int                               `json:"limit,omitempty"`
	Offset   int                               `json:"offset,omitempty"`
	Response string                            `json:"response,omitempty"`
	Total    int                               `json:"total,omitempty"`
}

type MangaDexSearchChapterResultData struct {
	MangaDexResultItem
	Relationships []MangaDexSearchResultDataRelationship `json:"relationships,omitempty"`
	Attributes    struct {
		Title              string `json:"title,omitempty"`
		Volume             string `json:"volume,omitempty"`
		Chapter            string `json:"chapter,omitempty"`
		CreatedAt          string `json:"created_at,omitempty"`
		Pages              int    `json:"pages,omitempty"`
		TranslatedLanguage string `json:"translated_language,omitempty"`
	} `json:"attributes,omitempty"`
}

type ChapterPagesResult struct {
	BaseURL string            `json:"baseUrl,omitempty"`
	Chapter ChapterPageResult `json:"chapter,omitempty"`
}

type ChapterPageResult struct {
	Hash string   `json:"hash,omitempty"`
	Data []string `json:"data,omitempty"`
}
