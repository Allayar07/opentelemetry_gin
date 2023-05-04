package services

import (
	"practice_optelem/first-service/internal/redis_cache"
	"practice_optelem/first-service/internal/repositroy"
)

type Service struct {
	File *FileService
}

func NewService(repo *repositroy.Repository, caches *redis_cache.Cache) *Service {
	return &Service{
		File: NewFileService(repo.File, caches),
	}
}
