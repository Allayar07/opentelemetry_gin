package repositroy

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.opentelemetry.io/otel"
)

type FileRepos struct {
	db *pgxpool.Pool
}

func NewFileRepos(db *pgxpool.Pool) *FileRepos {
	return &FileRepos{
		db: db,
	}
}

func (r *FileRepos) AddFile(ctx context.Context, name string, size int) error {
	ctxt, span := otel.Tracer("practice-service").Start(ctx, "Repository.AddFile")
	defer span.End()

	query := fmt.Sprintf("INSERT INTO files (name, size) values ($1, $2)")
	_, err := r.db.Query(ctxt, query, name, size)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}
