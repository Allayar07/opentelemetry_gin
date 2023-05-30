package handler

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"practice_optelem/first-service/internal/services"
)

type Handler struct {
	Service *services.Service
	Tracer  bool
}

func NewHandler(service *services.Service, traceOn bool) *Handler {
	return &Handler{
		Service: service,
		Tracer:  traceOn,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	app := gin.Default()
	// connect tracer with handlers
	app.Use(otelgin.Middleware("1-service"))
	app.POST("/add/:name", h.AddFile)
	app.POST("/set", h.SetHash)
	app.GET("/call-service", h.CallSecondService)

	return app
}
