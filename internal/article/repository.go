package article

import (
	"context"
	"gorm.io/gorm"
)

type Repository interface {
	CreateArticle(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, article *Article) error
	FindByID(ctx context.Context, id uint) (*Article, error)
	FindAll(ctx context.Context) ([]*Article, error)
	DeleteArticle(ctx context.Context, article *Article, id uint) error
}

type GormRepository struct {
	*gorm.DB
}

func NewRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{DB: db}
}

func (r *GormRepository) CreateArticle(ctx context.Context, article *Article) error {
	err := r.DB.WithContext(ctx).Create(article).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) UpdateArticle(ctx context.Context, article *Article) error {
	err := r.DB.WithContext(ctx).Save(article).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) FindByID(ctx context.Context, id uint) (*Article, error) {

	var article Article

	err := r.DB.WithContext(ctx).Preload("Author").First(&article, id).Error

	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *GormRepository) FindAll(ctx context.Context) ([]*Article, error) {
	var articles []*Article

	err := r.DB.WithContext(ctx).Preload("Author").Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *GormRepository) DeleteArticle(ctx context.Context, article *Article, id uint) error {

	result := r.DB.WithContext(ctx).Delete(article, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
