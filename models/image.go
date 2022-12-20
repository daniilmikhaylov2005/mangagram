package models

type Image struct {
	ID     int    `json:"id"`
	Url    string `json:"url"`
	Author string `json:"author"`
}
