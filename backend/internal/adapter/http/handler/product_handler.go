package handler

import (
	"ecommerce/internal/adapter/http/middleware"
	"ecommerce/internal/core/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *product.Service
}

func NewHandler(s *product.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register(r *gin.Engine) {
	group := r.Group("/products")

	group.Use(middleware.AuthMiddleare())
	{
		group.GET("/", h.getAll)
		group.POST("/", h.create)
		group.GET("/:id", h.getOne)
	}
}

func (h *Handler) getAll(c *gin.Context) {
	products, err := h.service.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) create(c *gin.Context) {
	product := product.Product{}
	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Create(product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error to save product",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": "product created",
	})
}

func (h *Handler) getOne(c *gin.Context) {
	productID := c.Params[0].Value

	intProductId, err := strconv.Atoi(productID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Product not found or not exists",
			"Error":   err.Error(),
		})
		return
	}

	product, err := h.service.GetOne(intProductId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Product not found or not exists",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
