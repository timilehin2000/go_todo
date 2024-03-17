package models

type Todo struct {
	Id          int    `json:"id"`
	Item        string `json:"item"`
	IsCompleted bool   `json:"is_completed"`
}
