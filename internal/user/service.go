package user

import (
	"context"
	"github.com/Bobby-P-dev/go-clean-api/pkg/bcrypt"
	"github.com/Bobby-P-dev/go-clean-api/pkg/customerr"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateUser(ctx context.Context, r CreateUserRequest) (*UserResponse, error) {

	pwd, err := bcrypt.HashPwd(r.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: r.Username,
		Email:    r.Email,
		Password: pwd,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *Service) ListUsers(ctx context.Context, page, limit int) ([]*UserResponse, int64, error) {

	users, total, err := s.repository.ListUsers(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var userResponses []*UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return userResponses, total, nil
}

func (s *Service) LoginUser(ctx context.Context, email, password string) (*UserResponseLogin, error) {

	user, err := s.repository.LoginUser(ctx, email)
	if err != nil {
		return nil, err
	}

	match := bcrypt.CheckPwdHash(password, user.Password)

	if !match {
		return nil, customerr.ErrUnauthorized
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_KEY")

	if secret == "" {
		return nil, customerr.ErrInternal
	}

	key := []byte(secret)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	res := &UserResponseLogin{
		Success: true,
		Message: "login successful yeah",
		Jwt:     Jwt{Token: tokenString},
	}
	return res, nil
}
