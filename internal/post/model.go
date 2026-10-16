package post

import (
	"errors"
	"strings"
	"time"
)

// Post represents the post in the blog
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}

// Prepare use validate and format functions
func (p *Post) Prepare() error {
	p.format()
	if erro := p.validate(); erro != nil {
		return erro
	}

	return nil
}

// validate if Title and Content is not empty
func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("failed to create: title is required")
	}
	if p.Content == "" {
		return errors.New("failed to create: content is required")
	}
	return nil
}

// format the post title
func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
}
