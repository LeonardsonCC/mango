package entity

type MangaDexSearchResult struct {
	Data     []MangaDexSearchResultData `json:"data,omitempty"`
	Limit    int                        `json:"limit,omitempty"`
	Offset   int                        `json:"offset,omitempty"`
	Response string                     `json:"response,omitempty"`
	Total    int                        `json:"total,omitempty"`
}

type MangaDexSearchResultData struct {
	MangaDexResultItem
	Relationships []MangaDexSearchResultDataRelationship `json:"relationships,omitempty"`
	Attributes    struct {
		Title MangaDexResultLanguages `json:"title,omitempty"`
	} `json:"attributes,omitempty"`
}

type MangaDexSearchResultDataRelationship struct {
	MangaDexResultItem
	Attributes struct {
		FileName    string `json:"fileName,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"attributes,omitempty"`
}
