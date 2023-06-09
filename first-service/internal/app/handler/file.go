package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
)

func (h *Handler) AddFile(c *gin.Context) {
	var (
		ctx  context.Context
		span trace.Span
	)
	name := c.Query("name")
	if h.Tracing {
		ctx, span = otel.Tracer("first-service").Start(c.Request.Context(), "Delivery.AddFile")
		defer span.End()
	} else {
		ctx = context.Background()
	}

	if err := h.Service.File.Add(ctx, name, 170); err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, "OK")
}

func (h *Handler) SetHash(c *gin.Context) {
	var (
		ctx  context.Context
		span trace.Span
	)
	if h.Tracing {
		ctx, span = otel.Tracer("1-service").Start(c.Request.Context(), "Delivery.SetHash")
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

func (h *Handler) CallSecondService(c *gin.Context) {
	client := http.Client{}
	var (
		req *http.Request
		err error
	)
	if h.Tracing {
		ctx, span := otel.Tracer("1-service").Start(c.Request.Context(), "1-service-handler")
		defer span.End()
		span.RecordError(errors.New("hello world"))
		req, err = http.NewRequestWithContext(ctx, "GET", "http://service-2:8081/second-service/say_hello", nil)
		span.AddEvent("Errors:", trace.WithAttributes(
			attribute.String("log.errors", fmt.Sprintf("%s", "oops error here appeared")),
		))
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	} else {
		req, err = http.NewRequestWithContext(c.Request.Context(), "GET", "http://localhost:8081/second-service/say_hello", nil)
		if err != nil {

			c.JSON(500, err.Error())
			return
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	result, err := io.ReadAll(resp.Body)
	_, err = c.Writer.Write(result)
	if err != nil {
		c.JSON(500, err)
		return
	}
}
