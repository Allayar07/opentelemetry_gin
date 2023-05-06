package handler

import (
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
	name := c.Query("name")
	ctx, span := otel.Tracer("first-service").Start(c.Request.Context(), "Delivery.AddFile")
	defer span.End()
	if err := h.Service.File.Add(ctx, name, 170); err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, "OK")
}

func (h *Handler) SetHash(c *gin.Context) {
	ctx, span := otel.Tracer("first-service").Start(c.Request.Context(), "Delivery.SetHash")
	defer span.End()
	err := h.Service.File.SetHash(ctx)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "OK")
}

func (h *Handler) CallSecondService(c *gin.Context) {
	client := http.Client{}
	//name := c.Param("name")
	ctx, span := otel.Tracer("1-service").Start(c.Request.Context(), "1-service-handler")
	defer span.End()
	span.RecordError(errors.New("hello world"))
	req, err := http.NewRequestWithContext(ctx, "GET", "http://service-2:8081/second-service/say_hello", nil)
	if err != nil {
		span.AddEvent("Errors:", trace.WithAttributes(
			attribute.String("log.errors", fmt.Sprintf("%s", err)),
		))
		c.JSON(500, err.Error())
		return
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
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
