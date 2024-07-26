package entity

type MangaDexResultItem struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type MangaDexResultLanguages struct {
	JA string `json:"ja,omitempty"`
	EN string `json:"en,omitempty"`
}
