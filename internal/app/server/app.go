package server

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"log"
	"practice_optelem/internal/app/handler"
	"practice_optelem/internal/redis_cache"
	"practice_optelem/internal/repositroy"
	"practice_optelem/internal/services"
	"practice_optelem/pkg"
)

func Init() {
	db, err := repositroy.NewPostgresDB(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	repos := repositroy.NewRepository(db)
	client, err := pkg.NewRedisClient(7, context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	cacheService := redis_cache.NewCache(client)
	serv := services.NewService(repos, cacheService)
	handlers := handler.NewHandler(serv)

	srv := new(Server)
	//initializing trace provider
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	defer func() {
		if err = tp.Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}()

	if err = srv.Run("127.0.0.1:8080", handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}

}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("practice-service"),
			//attribute.String("environment", environment),
			//attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}
