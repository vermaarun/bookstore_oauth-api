package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vermaarun/bookstore_oauth-api/src/domain/access_token"
	"github.com/vermaarun/bookstore_oauth-api/src/utils/errors"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var token access_token.AccessToken
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := h.service.Create(token); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, token)
}