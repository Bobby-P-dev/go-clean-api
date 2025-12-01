package user

import (
	"net/http"

	"github.com/Bobby-P-dev/go-clean-api.git/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User Info"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]interface{}
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	user := UserResponse{
		ID:       "1",
		Username: req.Username,
		Email:    req.Email,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user created",
		"data":    user,
	})
}
