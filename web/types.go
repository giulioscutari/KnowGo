package main

import "strings"

type Document struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	Paragraphs []string `json:"paragraphs"`
}

func NewDocument(id int, title, body string) *Document {
	return &Document{
		Id:         id,
		Title:      title,
		Body:       body,
		Paragraphs: strings.Split(body, "."),
	}
}
