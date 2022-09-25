package models

type Quote struct {
	Id         string `json:"id"`
	Quote      string `json:"quote"`
	AuthorName string `json:"author_name"`
}

type Status struct {
	Status string `json:"status"`
}
