package upload

import (
	"context"
	uploadDomain "github.com/aryahmph/restaurant/internal/domain/upload"
)

type Repository interface {
	Insert(ctx context.Context, filename string) (id string, err error)
	FindByID(ctx context.Context, id string) (upload uploadDomain.Upload, err error)
	Delete(ctx context.Context, id string) (rid string, err error)
}
