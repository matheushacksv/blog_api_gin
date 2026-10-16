package post

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// post represents a post repository
type post struct {
	db *pgxpool.Pool
}

// NewRepositoryPost create a new post repository
func NewRepositoryPost(db *pgxpool.Pool) *post {
	return &post{db}
}

// CreatePost adds a post in database
func (p post) CreatePost(ctx context.Context, post Post) (uint64, error) {
	erro := p.db.QueryRow(
		ctx,
		"INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING id, created_at, updated_at",
		post.Title, post.Content,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if erro != nil {
		return 0, erro
	}
	return post.ID, nil
}

// SearchPosts search by title from database
func (p post) SearchPosts(ctx context.Context, title string) ([]Post, error) {
	rows, erro := p.db.Query(
		ctx,
		"SELECT id, title, content, created_at, updated_at FROM posts WHERE title ILIKE '%' || $1 || '%' ORDER BY created_at",
		title,
	)
	if erro != nil {
		return []Post{}, erro
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post

		if erro := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}
	return posts, nil
}

// GetAPost get a post by id from the database
func (p post) GetAPost(ctx context.Context, ID uint64) (Post, error) {
	var post Post
	erro := p.db.QueryRow(
		ctx,
		"SELECT id, title, content, created_at, updated_at FROM posts WHERE id = $1",
		ID,
	).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if erro != nil {
		return Post{}, erro
	}
	return post, nil
}

// UpdatePost updates a data from the database
func (p post) UpdatePost(ctx context.Context, ID uint64, title, content string) error {
	_, erro := p.db.Exec(
		ctx,
		"UPDATE posts SET title = $1, content = $2, updated_at = now() WHERE id = $3",
		title, content, ID,
	)
	if erro != nil {
		return erro
	}

	return nil
}

// DeletePost deletes from the database
func (p post) DeletePost(ctx context.Context, ID uint64) error {
	_, erro := p.db.Exec(
		ctx,
		"DELETE FROM posts WHERE id = $1",
		ID,
	)
	if erro != nil {
		return erro
	}
	return nil
}
