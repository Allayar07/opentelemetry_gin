package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func (h *Handler) AddFile(c *gin.Context) {
	name := c.Query("name")
	var (
		ctx  context.Context
		span trace.Span
	)
	if h.Tracing {
		ctx, span = otel.Tracer("practice-service").Start(c.Request.Context(), "Delivery.AddFile")
		defer span.End()
	} else {
		ctx = context.Background()
	}

	if err := h.Service.File.Add(ctx, name, 170); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "OK. added to redis!")
}

func (h *Handler) SetHash(c *gin.Context) {
	var (
		ctx  context.Context
		span trace.Span
	)
	if h.Tracing {
		ctx, span = otel.Tracer("practice-service").Start(c.Request.Context(), "Delivery.SetHash")
		defer span.End()
	} else {
		ctx = context.Background()
	}

	err := h.Service.File.SetHash(ctx)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "OK")
}

func (h *Handler) SayHello(c *gin.Context) {
	if h.Tracing {
		_, span := otel.Tracer("2-service").Start(c.Request.Context(), "2-services-handler")
		otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		defer span.End()
	}

	_, err := c.Writer.Write([]byte("Hello every one!!!"))
	if err != nil {
		c.JSON(500, err)
		return
	}
}
