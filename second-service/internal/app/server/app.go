package server

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"log"
	"practice_optelem/second-service/internal/app/handler"
	"practice_optelem/second-service/internal/redis_cache"
	repositroy2 "practice_optelem/second-service/internal/repositroy"
	"practice_optelem/second-service/internal/services"
)

func Init(port string) {
	//db, err := repositroy2.NewPostgresDB(context.Background())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	traceOn := viper.GetBool("service.trace_on")
	repos := repositroy2.NewRepository(&pgxpool.Pool{})
	//client, err := pkg.NewRedisClient(7, context.Background())
	//if err != nil {
	//	log.Fatalln(err)
	//}

	cacheService := redis_cache.NewCache(&redis.Client{})
	serv := services.NewService(repos, cacheService)
	handlers := handler.NewHandler(serv, traceOn)

	srv := new(Server)
	if traceOn {
		//initializing trace provider
		tracePrv, err := tracerProvider("http://zipkin:9411/api/v2/spans")
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if err = tracePrv.Shutdown(context.Background()); err != nil {
				log.Fatalln(err)
			}
		}()
	} else {
		logrus.Info("tracer off!!!")
	}

	if err := srv.Run(":"+port, handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}

}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	//exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	//if err != nil {
	//	return nil, err
	//}
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	exporter, err := zipkin.New(
		url,
	)
	if err != nil {
		return nil, err
	}
	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("2-service"),
		)),
	)
	otel.SetTracerProvider(traceProvider)
	return traceProvider, nil
}

// InitConfig initializes configuration.
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return errors.Wrap(viper.ReadInConfig(), "can not read config")
}
