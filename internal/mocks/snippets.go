package mocks

import (
	"time"

	"github.com/sam-maton/snippetbox/internal/models"
)

var MockSnippet = models.Snippet{
	ID:      1,
	Title:   "Test snippet 1",
	Content: "Test snippet content",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (models.Snippet, error) {
	switch id {
	case 1:
		return MockSnippet, nil
	default:
		return models.Snippet{}, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]models.Snippet, error) {
	return []models.Snippet{MockSnippet}, nil
}
