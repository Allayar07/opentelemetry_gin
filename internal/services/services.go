package services

import (
	"practice_optelem/internal/redis_cache"
	"practice_optelem/internal/repositroy"
)

type Service struct {
	File *FileService
}

func NewService(repo *repositroy.Repository, caches *redis_cache.Cache) *Service {
	return &Service{
		File: NewFileService(repo.File, caches),
	}
}
