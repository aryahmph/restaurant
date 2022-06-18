package category

import (
	"context"
	categoryDomain "github.com/aryahmph/restaurant/internal/domain/category"
)

type Repository interface {
	Insert(ctx context.Context, category categoryDomain.Category) (id string, err error)
	FindAll(ctx context.Context) (categories []categoryDomain.Category, err error)
	FindByID(ctx context.Context, id string) (category categoryDomain.Category, err error)
	FindByName(ctx context.Context, name string) (category categoryDomain.Category, err error)
	Update(ctx context.Context, category categoryDomain.Category) (id string, err error)
	Delete(ctx context.Context, id string) (rid string, err error)
}
