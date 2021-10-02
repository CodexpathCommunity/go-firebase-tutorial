package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID          uint32    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (post *Post) Prepare() {
	post.ID = 0
	post.Title = html.EscapeString(strings.TrimSpace(post.Title))
	post.Description = html.EscapeString(strings.TrimSpace(post.Description))
	post.CreatedAt = time.Now()
}

func (post *Post) Validate() error {
	if post.Title == "" {
		return errors.New("Title required")
	}
	if post.Description == "" {
		return errors.New("Description required")
	}
	return nil
}
