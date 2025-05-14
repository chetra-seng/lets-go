package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

type SnippetModelInterface interface {
  Insert(title string, content string, expires int) (int, error)
  Get(id int) (Snippet, error)
  Latest() ([]Snippet, error)
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUE (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := "SELECT id, title, content, created, expires from snippets where expires > UTC_TIMESTAMP() AND id = ?"
	var s Snippet

	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	// No need to defer connect since it only contain a single row
	// If it can't be read, then the connection will be close

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	stmt := "SELECT id, title, content, created, expires from snippets where expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10"
	var snippets []Snippet

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	// Defer connection close in case sometime goes wrong
	// Because there are multiple rows, we can't tell which row will fail
	defer rows.Close()

	for rows.Next() {
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	// Check errors after the loop or look for any error when iterating
	if err = rows.Err(); err != nil {
    return nil, err
	}

	// No errors encountered
	return snippets, nil
}
