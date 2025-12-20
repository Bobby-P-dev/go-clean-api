package article

import (
	"github.com/Bobby-P-dev/go-clean-api/pkg/response"
	validator2 "github.com/Bobby-P-dev/go-clean-api/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateArticle(c *gin.Context) {

	ctx := c.Request.Context()
	var req CreateArticleRequest

	rawId, ok := c.Get("user_id")

	if !ok {
		response.Error(c, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	var userID uint
	if val, ok := rawId.(float64); ok {
		userID = uint(val)
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		valError := validator2.FormatValidationError(err)
		response.Error(c, http.StatusBadRequest, "invalid request body", valError)
		return
	}

	article, err := h.service.CreateArticle(ctx, &req, userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Created(c, "article created successfully", article)
}
