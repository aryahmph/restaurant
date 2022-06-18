package category

import (
	"context"
	categoryDomain "github.com/aryahmph/restaurant/internal/domain/category"
	userDomain "github.com/aryahmph/restaurant/internal/domain/user"
)

type Usecase interface {
	Create(ctx context.Context, category categoryDomain.Category, adminID string) (id string, err error)
	List(ctx context.Context) (categories []categoryDomain.Category, err error)
	Update(ctx context.Context, category categoryDomain.Category, adminID string) (id string, err error)
	Delete(ctx context.Context, categoryID string, adminID string) (rid string, err error)

	checkAdminByID(ctx context.Context, id string) (user userDomain.User, err error)
	nameAlreadyExist(ctx context.Context, username string) (err error)
}
