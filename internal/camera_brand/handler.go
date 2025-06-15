package camera_brand

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/brands")
	group.POST("/", h.Create)
	group.GET("/", h.GetAll)
	group.GET("/:id", h.GetOne)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}

func (h *Handler) Create(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&req); err != nil || req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	brand, err := h.service.CreateCameraBrand(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, brand)
}

func (h *Handler) GetAll(c *gin.Context) {
	brands, err := h.service.GetAllCameraBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get all failed"})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (h *Handler) GetOne(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	brand, err := h.service.GetCameraBrandByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&req); err != nil || req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := h.service.UpdateCameraBrand(uint(id), req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteCameraBrand(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
