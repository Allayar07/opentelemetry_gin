package handler

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func (h *Handler) AddFile(c *gin.Context) {
	name := c.Query("name")
	ctx, span := otel.Tracer("practice-service").Start(c.Request.Context(), "Delivery.AddFile")
	defer span.End()
	if err := h.Service.File.Add(ctx, name, 170); err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, "OK")
}

func (h *Handler) SetHash(c *gin.Context) {
	ctx, span := otel.Tracer("practice-service").Start(c.Request.Context(), "Delivery.SetHash")
	defer span.End()
	err := h.Service.File.SetHash(ctx)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "OK")
}

func (h *Handler) SayHello(c *gin.Context) {
	_, span := otel.Tracer("2-service").Start(c.Request.Context(), "2-services-handler")
	otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
	defer span.End()
	_, err := c.Writer.Write([]byte("Hello every one!!!"))
	if err != nil {
		c.JSON(500, err)
		return
	}
}
