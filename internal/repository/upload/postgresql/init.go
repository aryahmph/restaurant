package postgresql

import (
	"context"
	"database/sql"
	"errors"
	errorCommon "github.com/aryahmph/restaurant/common/error"
	uploadDomain "github.com/aryahmph/restaurant/internal/domain/upload"
)

type postgreSqlUploadRepositoryImpl struct {
	db *sql.DB
}

func NewPostgreSqlUploadRepositoryImpl(db *sql.DB) postgreSqlUploadRepositoryImpl {
	return postgreSqlUploadRepositoryImpl{db: db}
}

func (r postgreSqlUploadRepositoryImpl) Insert(ctx context.Context, filename string) (id string, err error) {
	row := r.db.QueryRowContext(ctx, "INSERT INTO uploads(filename) VALUES ($1) RETURNING id;", filename)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("file not found")
	}
	return id, err
}

func (r postgreSqlUploadRepositoryImpl) FindByID(ctx context.Context, id string) (upload uploadDomain.Upload, err error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, filename, created_at, updated_at FROM uploads WHERE id = $1 LIMIT 1;", id)

	err = row.Scan(&upload.ID, &upload.Filename, &upload.CreatedAt, &upload.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return upload, errorCommon.NewNotFoundError("file not found")
	}
	return upload, err
}

func (r postgreSqlUploadRepositoryImpl) Delete(ctx context.Context, id string) (rid string, err error) {
	row := r.db.QueryRowContext(ctx, "DELETE FROM uploads WHERE id=$1 RETURNING id;", id)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return rid, errorCommon.NewNotFoundError("file not found")
	}
	return rid, err
}
