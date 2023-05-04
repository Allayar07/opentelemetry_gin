package handler

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"practice_optelem/second-service/internal/services"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.Default()
	// connect tracer with handlers
	app.Use(otelgin.Middleware("2-service"))
	app.POST("/second-service/add/:name", h.AddFile)
	app.POST("/set", h.SetHash)
	app.GET("/second-service/say_hello", h.SayHello)

	return app
}