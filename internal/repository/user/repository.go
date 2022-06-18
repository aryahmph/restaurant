package user

import (
	"context"
	userDomain "github.com/aryahmph/restaurant/internal/domain/user"
)

type Repository interface {
	Insert(ctx context.Context, user userDomain.User) (id string, err error)
	FindAll(ctx context.Context) (users []userDomain.User, err error)
	FindByID(ctx context.Context, id string) (user userDomain.User, err error)
	FindByUsername(ctx context.Context, username string) (user userDomain.User, err error)
	Update(ctx context.Context, user userDomain.User) (id string, err error)
	Delete(ctx context.Context, id string) (rid string, err error)
}
