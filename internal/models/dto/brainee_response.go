package dto

type BraineeResponse struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
	Brand  string `json:"brand"`
}
