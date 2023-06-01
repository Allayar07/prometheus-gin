package handler

import (
	"gin_prometheus/internal/app/metrics"
	"gin_prometheus/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.Default()
	p := metrics.NewPrometheus()
	p.Use(app)
	app.POST("/create", h.Create)
	app.GET("/say", h.SaySomething)
	app.GET("/just", h.JustDoIt)
	return app
}
