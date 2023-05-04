package services

import (
	"context"
	"go.opentelemetry.io/otel"
	"practice_optelem/second-service/internal/redis_cache"
	"practice_optelem/second-service/internal/repositroy"
)

type FileService struct {
	repo  *repositroy.FileRepos
	cache *redis_cache.Cache
}

func NewFileService(repo *repositroy.FileRepos, cache *redis_cache.Cache) *FileService {
	return &FileService{
		repo:  repo,
		cache: cache,
	}
}

func (s *FileService) Add(ctx context.Context, name string, size int) error {
	ctxt, span := otel.Tracer("practice-service").Start(ctx, "Service.AddFile")
	defer span.End()

	return s.repo.AddFile(ctxt, name, size)
}

func (s *FileService) SetHash(ctx context.Context) error {
	ctxt, span := otel.Tracer("practice-service").Start(ctx, "Service.SetHash")
	defer span.End()
	return s.cache.SetHash(ctxt)
}
