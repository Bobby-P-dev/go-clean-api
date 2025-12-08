package user

import (
	"errors"
	"github.com/Bobby-P-dev/go-clean-api/pkg/customerr"
	validator2 "github.com/Bobby-P-dev/go-clean-api/pkg/validator"
	"net/http"

	"github.com/Bobby-P-dev/go-clean-api/pkg/response"
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
		valErrors := validator2.FormatValidationError(err)
		response.Error(c, http.StatusBadRequest, "invalid request body", valErrors)
		return
	}

	user, err := h.service.CreateUser(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, customerr.ErrConflict):
			response.Error(c, http.StatusConflict, err.Error(), nil)
		default:
			response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}
	response.Created(c, "user created successfully", user)
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
		valErrors := validator2.FormatValidationError(err)
		response.Error(c, http.StatusBadRequest, "invalid query parameters", valErrors)
		return
	}

	users, total, err := h.service.ListUsers(ctx, q.Page, q.Limit)
	if err != nil {
		switch {
		case errors.Is(err, customerr.ErrInternal):
			response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		default:
			response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	meta := &PaginationMeta{
		Page:  q.Page,
		Limit: q.Limit,
		Total: int(total),
	}
	response.SuccsesMeta(c, "Successfully retrieved users", users, meta)
}

// LoginUser godoc
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "User Credentials"
// @Success 200 {object} UserResponseLogin
// @Failure 401 {object} map[string]interface{}
// @Router /users/login [post]
func (h *Handler) LoginUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		valErrors := validator2.FormatValidationError(err)
		response.Error(c, http.StatusBadRequest, "invalid request body", valErrors)
		return
	}

	loginResp, err := h.service.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, customerr.ErrUnauthorized):
			response.Error(c, http.StatusUnauthorized, "invalid email or password", nil)
		default:
			response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	response.Success(c, "login successful", loginResp)
}
