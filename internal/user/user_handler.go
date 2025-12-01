package user

import (
	"net/http"

	"github.com/Bobby-P-dev/go-clean-api.git/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
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

	ctx := c.Request.Context()
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.service.CreateUser(ctx, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user created",
		"data":    user,
	})
}

// ListUsers godoc
// @Summary List users with pagination
// @Description Retrieve a paginated list of users
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of users per page"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /users [get]
func (h *Handler) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()

	var q PaginationQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid query parameters")
		return
	}

	users, total, err := h.service.ListUsers(ctx, q.Page, q.Limit)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "users retrieved",
		"data":    users,
		"meta": gin.H{
			"page":  q.Page,
			"limit": q.Limit,
			"total": total,
		},
	})
}
