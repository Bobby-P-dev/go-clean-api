package user

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ListUsers(page, limit int) ([]UserResponse, int) {
	users := []UserResponse{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
		{ID: 3, Username: "user3", Email: "user3@example.com"},
		{ID: 4, Username: "user4", Email: "user4@example.com"},
		{ID: 5, Username: "user5", Email: "user5@example.com"},
	}

	total := len(users)
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	start := (page - 1) * limit
	if start > total {
		return []UserResponse{}, total
	}

	end := start + limit
	if end > total {
		end = total
	}

	return users[start:end], total
}
