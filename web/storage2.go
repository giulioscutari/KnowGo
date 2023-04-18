package main

import (
	"fmt"

	"github.com/lib/pq"
)

func (s *PostgresStore) CreateDocument(d *Document) error {
	sqlStatement := `INSERT INTO document (title, body, paragraphs, created_at)
	VALUES
	($1, $2, $3, $4)`
	resp, err := s.db.Exec(
		sqlStatement,
		d.Title,
		d.Body,
		pq.Array(d.Paragraphs),
		d.CreatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("Query supposedly went well")
	}

	fmt.Printf("this is resp:\n%v\n", resp)

	return nil
}
