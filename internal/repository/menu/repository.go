package menu

import (
	"context"
	menuDomain "github.com/aryahmph/restaurant/internal/domain/menu"
)

type Repository interface {
	Insert(ctx context.Context, menu menuDomain.Menu) (id string, err error)
	FindAll(ctx context.Context) (menus []menuDomain.Menu, err error)
	FindByID(ctx context.Context, id string) (menu menuDomain.Menu, err error)
	FindByName(ctx context.Context, name string) (menu menuDomain.Menu, err error)
	Update(ctx context.Context, menu menuDomain.Menu) (id string, err error)
	Delete(ctx context.Context, menuID string) (rid string, err error)
}
