package handler

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"practice_optelem/first-service/internal/services"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//appRelic, err := newrelic.NewApplication(
	//	newrelic.ConfigAppName("practice app"),
	//	newrelic.ConfigLicense("here need license, but I can not get it :( "),
	//	newrelic.ConfigDebugLogger(os.Stdout),
	//	newrelic.ConfigCodeLevelMetricsEnabled(true),
	//)
	//if nil != err {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	app := gin.Default()
	// connect tracer with handlers
	app.Use(otelgin.Middleware("1-service") /* nrgin.Middleware(appRelic)*/)
	app.POST("/add/:name", h.AddFile)
	app.POST("/set", h.SetHash)
	app.GET("/call-service", h.CallSecondService)

	return app
}
