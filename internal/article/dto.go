package article

import "time"

type CreateArticleRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	AuthorID uint   `json:"author_id" binding:"required"`
}

type UpdateArticleRequest struct {
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
	AuthorID uint   `json:"author_id,omitempty"`
}

type ArticleResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	AuthorID    uint       `json:"author_id"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}
