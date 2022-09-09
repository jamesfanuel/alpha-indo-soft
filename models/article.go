package models

import "time"

type Article struct {
	ID uint
	Author string `json: "body"`
	Title string `json: "title"`
	Body string `json: "body"`
	CreatedAt time.Time
}