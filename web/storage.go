package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateDocument(*Document) error
	DeleteDocument(int) error
	UpdateDocument(*Document) error
	GetDocuments() ([]*Document, error)
	GetDocumentByID(int) (*Document, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_DB")
	connStr := "postgres://" + user + ":" + password + "@db/" + db_name + "?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil

}

func (s *PostgresStore) Init() error {
	return s.createDocumentTable()
}
func (s *PostgresStore) createDocumentTable() error {
	query := `CREATE TABLE IF NOT EXISTS document (
		id serial primary key,
		title varchar(255),
		body text,
		paragraphs text[],
		created_at timestamp

	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) DeleteDocument(id int) error {
	return nil
}
func (s *PostgresStore) UpdateDocument(*Document) error {
	return nil
}
func (s *PostgresStore) GetDocuments() ([]*Document, error) {
	rows, err := s.db.Query("select id, title, body, created_at from document")
	fmt.Println("we're inside GetDocuments")
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	docs := []*Document{}
	for rows.Next() {
		doc := new(Document)
		err := rows.Scan(
			&doc.Id,
			&doc.Title,
			&doc.Body,
			&doc.CreatedAt,
		)
		if err != nil {
			return nil, err

		}
		fmt.Println(doc)
		docs = append(docs, doc)

	}

	return docs, nil
}
func (s *PostgresStore) GetDocumentByID(id int) (*Document, error) {
	return nil, nil
}
