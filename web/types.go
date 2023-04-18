package main

import (
	"strings"
	"time"
)

type CreateDocumentRequest struct {
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	Paragraphs []string `json:"paragraphs"`
}

type Document struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Paragraphs []string  `json:"paragraphs"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewDocument(title, body string) *Document {
	return &Document{
		// Id:         rand.Intn(10000),
		Title:      title,
		Body:       body,
		Paragraphs: strings.Split(body, "."),
		CreatedAt:  time.Now().UTC(),
	}
}
