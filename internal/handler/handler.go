package handler

import (
	"github.com/gin-gonic/gin"
	"l0-project/internal/cache"
	"net/http"
)

type Handler struct {
	Router *gin.Engine
}

func (h *Handler) InitRoutes(c *cache.Cache) {
	h.Router = gin.Default()

	orders := h.Router.Group("/api/orders")
	{
		orders.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, c.GetAllOrders())
		})
		orders.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(http.StatusOK, c.GetOrderByID(id))
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
