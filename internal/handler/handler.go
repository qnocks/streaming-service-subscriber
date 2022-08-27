package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"l0-project/internal/cache"
	"net/http"
	"time"
)

type Handler struct {
	Router *gin.Engine
}

type errorResponse struct {
	Status    int64  `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func (h *Handler) InitRoutes(c *cache.Cache) {
	h.Router = gin.Default()

	orders := h.Router.Group("/api/orders")
	{
		orders.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			order := c.GetOrderByID(id)

			if order == nil {
				ctx.JSON(http.StatusNotFound, errorResponse{
					Status:    http.StatusNotFound,
					Error:     http.StatusText(http.StatusNotFound),
					Message:   fmt.Sprintf("Order with id: %s not found", id),
					Timestamp: time.Now().Unix(),
				})
			} else {
				ctx.JSON(http.StatusOK, order)
			}
		})
	}
}

func (h *Handler) InitUI() {
	h.Router.Static("/assets", "./assets")
	h.Router.LoadHTMLGlob("static/*.html")

	h.Router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
