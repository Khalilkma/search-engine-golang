package model

type Page struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Headings    string `json:"headings"`
	Content     string `json:"content"`
}
