package user

import (
	"context"
	"github.com/Bobby-P-dev/go-clean-api.git/pkg/customerr"
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	ListUsers(ctx context.Context, page, limit int) ([]*User, int64, error)
	LoginUser(ctx context.Context, email string) (*User, error)
}

type GormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) CreateUser(ctx context.Context, user *User) error {

	err := r.db.WithContext(ctx).Create(user).Error

	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return customerr.ErrConflict
		}
		return err
	}
	return nil
}

func (r *GormRepository) ListUsers(ctx context.Context, page, limit int) ([]*User, int64, error) {
	var users []*User
	var total int64

	q := r.db.WithContext(ctx).Model(&User{})

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, customerr.ErrInternal
	}

	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	if err := q.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, customerr.ErrInternal
	}
	return users, total, nil
}

func (r *GormRepository) LoginUser(ctx context.Context, email string) (*User, error) {

	var user User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, customerr.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
