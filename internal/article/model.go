package article

import (
	"github.com/Bobby-P-dev/go-clean-api/internal/user"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title       string     `gorm:"size:255;not null;" json:"title"`
	Content     string     `gorm:"type:text;not null;" json:"content"`
	AuthorID    uint       `gorm:"not null;" json:"author_id"`
	Author      user.User  `gorm:"foreignKey:AuthorID;references:ID" json:"author"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}
