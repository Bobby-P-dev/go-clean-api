package user

import (
	"context"
	"github.com/Bobby-P-dev/go-clean-api/pkg/bcrypt"
	"github.com/Bobby-P-dev/go-clean-api/pkg/customerr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

type MocRepository struct {
	mock.Mock
}

var _ Repository = (*MocRepository)(nil)

func (m *MocRepository) CreateUser(ctx context.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MocRepository) LoginUser(ctx context.Context, email string) (*User, error) {
	args := m.Called(ctx, email)
	if u := args.Get(0); u != nil {
		return u.(*User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MocRepository) ListUsers(ctx context.Context, page, limit int) ([]*User, int64, error) {
	args := m.Called(ctx, page, limit)

	var users []*User
	if u := args.Get(0); u != nil {
		users = u.([]*User)
	}

	var total int64
	if t := args.Get(1); t != nil {
		total = t.(int64)
	}

	return users, total, args.Error(2)
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MocRepository)
	service := NewService(mockRepo)

	req := CreateUserRequest{
		Username: "testuser",
		Email:    "example@gmail.com",
		Password: "password123",
	}

	mockRepo.
		On("CreateUser", mock.Anything, mock.Anything).
		Return(nil)

	data, err := service.CreateUser(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, req.Username, data.Username)

	mockRepo.AssertExpectations(t)
}

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MocRepository)
	service := NewService(mockRepo)

	users := []*User{
		{Username: "user1", Email: "example1@gmail.com"},
		{Username: "user2", Email: "example2@gmail.com"},
	}
	total := int64(len(users))

	mockRepo.
		On("ListUsers", mock.Anything, 1, 10).
		Return(users, total, nil)

	resp, count, err := service.ListUsers(ctx, 1, 10)

	assert.NoError(t, err)
	assert.Equal(t, total, count)
	assert.Len(t, resp, 2)
	assert.Equal(t, "user1", resp[0].Username)
	assert.Equal(t, "user2", resp[1].Username)

	mockRepo.AssertExpectations(t)
}

func TestLoginUser_Success(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MocRepository)
	service := NewService(mockRepo)

	// siapkan password hash
	rawPassword := "password123"
	hashed, err := bcrypt.HashPwd(rawPassword)
	assert.NoError(t, err)

	// user dari "database"
	mockUser := &User{
		Username: "testuser",
		Email:    "example@gmail.com",
		Password: hashed,
	}

	// set JWT_KEY supaya tidak kosong
	_ = os.Setenv("JWT_KEY", "test-secret-key")

	// expectation: repo.LoginUser dipanggil dengan email yang benar
	mockRepo.
		On("LoginUser", mock.Anything, mockUser.Email).
		Return(mockUser, nil)

	// call service
	res, err := service.LoginUser(ctx, mockUser.Email, rawPassword)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Sucsses)
	assert.Equal(t, "login successful yeah", res.Message)
	assert.NotEmpty(t, res.Jwt.Token)

	mockRepo.AssertExpectations(t)
}

func TestLoginUser_WrongPassword(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MocRepository)
	service := NewService(mockRepo)

	rawPassword := "password123"
	hashed, err := bcrypt.HashPwd(rawPassword)
	assert.NoError(t, err)

	mockUser := &User{
		Username: "testuser",
		Email:    "example@gmail.com",
		Password: hashed,
	}

	// repo tetap mengembalikan user yang valid
	mockRepo.
		On("LoginUser", mock.Anything, mockUser.Email).
		Return(mockUser, nil)

	// tapi password yang dikirim salah
	res, err := service.LoginUser(ctx, mockUser.Email, "wrong-password")

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, customerr.ErrUnauthorized, err)

	mockRepo.AssertExpectations(t)
}
