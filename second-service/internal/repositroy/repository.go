package repositroy

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	File *FileRepos
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		File: NewFileRepos(db),
	}
}
