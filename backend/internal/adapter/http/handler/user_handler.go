package handler

import (
	"ecommerce/internal/adapter/http/cookie"
	"ecommerce/internal/core/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *user.Service
}

func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/login", h.login)
	r.POST("/user", h.create)
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email ou senha incorretos"})
		return
	}

	cookie.SetAuthCookie(c.Writer, token)

	c.JSON(http.StatusOK, gin.H{"message": "login efetuado"})
}

func (h *UserHandler) create(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	err := h.service.Create(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Erro ao registrar usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success!"})
}
