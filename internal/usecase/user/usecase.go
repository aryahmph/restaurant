package user

import (
	"context"
	userDomain "github.com/aryahmph/restaurant/internal/domain/user"
)

type Usecase interface {
	Register(ctx context.Context, user userDomain.User, adminID string) (id string, err error)
	List(ctx context.Context, adminID string) (users []userDomain.User, err error)
	GetByID(ctx context.Context, id string, adminID string) (user userDomain.User, err error)
	Update(ctx context.Context, user userDomain.User, adminID string) (id string, err error)
	Delete(ctx context.Context, id string, adminID string) (rid string, err error)

	checkAdminByID(ctx context.Context, id string) (admin userDomain.User, err error)
	checkSameUserByID(ctx context.Context, id string) (user userDomain.User, err error)
	usernameAlreadyExist(ctx context.Context, username string) (err error)
}
