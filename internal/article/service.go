package article

import (
	"context"
	"github.com/Bobby-P-dev/go-clean-api/pkg/customerr"
)

type ServiceInterfaces interface {
	CreateArticle(ctx context.Context, req *CreateArticleRequest, userID uint) (*ArticleResponse, error)
	UpdateArticle(ctx context.Context, id uint, req *UpdateArticleRequest) (*ArticleResponse, error)
	GetArticleByID(ctx context.Context, id uint) (*ArticleResponse, error)
	GetAllArticles(ctx context.Context) ([]*ArticleResponse, error)
	DeleteArticle(ctx context.Context, id uint) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateArticle(ctx context.Context, req *CreateArticleRequest, UserID uint) (*ArticleResponse, error) {

	article := &Article{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: UserID,
	}

	if UserID == 0 || req.Title == "" || req.Content == "" {
		return nil, customerr.ErrBadRequest
	}

	if err := s.repo.CreateArticle(ctx, article); err != nil {
		return nil, err
	}

	return &ArticleResponse{
		ID:       article.ID,
		Title:    article.Title,
		Content:  article.Content,
		AuthorID: article.AuthorID,
	}, nil
}
