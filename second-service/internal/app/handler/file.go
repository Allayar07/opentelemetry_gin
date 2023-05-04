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

//var propagator = propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
//
//func handleRequest(w http.ResponseWriter, r *http.Request) {
//	// Extract the trace context from the incoming request
//	ctx := r.Context()
//	spanContext := propagator.Extract(ctx, propagation.HeaderCarrier(r.Header))
//	var span trace.Span
//	if spanContext.IsValid() {
//		ctx, span = tracer.Start(ctx, "my-operation", trace.WithSpanContext(spanContext))
//	} else {
//		ctx, span = tracer.Start(ctx, "my-operation")
//	}
//	defer span.End()
//
//	// Do some work here...
//
//	// Propagate the trace context to the outgoing request
//	req, _ := http.NewRequest("GET", "http://remote-service:8080", nil)
//	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Printf("failed to make outgoing request: %v", err)
//	} else {
//		defer resp.Body.Close()
//		// Do something with the response...
//	}
//}
//
//func Start(tracer *trace.TracerProvider) {
//	// Start the server and instrument it with OpenTelemetry
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", handleRequest)
//	handler := otelhttp.NewHandler(mux, "server")
//	http.ListenAndServe(":8080", handler)
//}
