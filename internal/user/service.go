package user

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateUser(ctx context.Context, r CreateUserRequest) (*UserResponse, error) {
	user := &User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
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
