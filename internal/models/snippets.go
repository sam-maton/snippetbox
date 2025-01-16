package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Create  time.Time
	Expire  time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	var id int
	expireDate := time.Now().Add((time.Hour * 24) * time.Duration(expires))

	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, NOW(), $3) RETURNING id`

	err := m.DB.QueryRow(stmt, title, content, expireDate).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
