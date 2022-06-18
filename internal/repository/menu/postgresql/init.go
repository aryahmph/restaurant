package postgresql

import (
	"context"
	"database/sql"
	"errors"
	errorCommon "github.com/aryahmph/restaurant/common/error"
	menuDomain "github.com/aryahmph/restaurant/internal/domain/menu"
)

type postgreSqlMenuRepositoryImpl struct {
	db *sql.DB
}

func NewPostgreSqlMenuRepositoryImpl(db *sql.DB) postgreSqlMenuRepositoryImpl {
	return postgreSqlMenuRepositoryImpl{db: db}
}

func (r postgreSqlMenuRepositoryImpl) Insert(ctx context.Context, menu menuDomain.Menu) (id string, err error) {
	row := r.db.QueryRowContext(ctx, "INSERT INTO menus(name, price, category_name, image_filename) VALUES ($1, $2, $3, $4) RETURNING id;",
		menu.Name, menu.Price, menu.CategoryName, menu.ImageFilename)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("menu not found")
	}
	return id, err
}

func (r postgreSqlMenuRepositoryImpl) FindAll(ctx context.Context) (menus []menuDomain.Menu, err error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, name, price, category_name, image_filename, created_at, updated_at, deleted_at FROM menus;")
	if err != nil {
		return menus, err
	}
	defer rows.Close()
	for rows.Next() {
		var menu menuDomain.Menu
		if err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Price,
			&menu.CategoryName,
			&menu.ImageFilename,
			&menu.CreatedAt,
			&menu.UpdatedAt,
			&menu.DeletedAt,
		); err != nil {
			return menus, err
		}
		menus = append(menus, menu)
	}
	if err := rows.Err(); err != nil {
		return menus, err
	}
	return menus, nil
}

func (r postgreSqlMenuRepositoryImpl) FindByID(ctx context.Context, id string) (menu menuDomain.Menu, err error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, name, price, category_name, image_filename, created_at, updated_at, deleted_at FROM menus WHERE id = $1 LIMIT 1;",
		id)

	err = row.Scan(
		&menu.ID,
		&menu.Name,
		&menu.Price,
		&menu.CategoryName,
		&menu.ImageFilename,
		&menu.CreatedAt,
		&menu.UpdatedAt,
		&menu.DeletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return menu, errorCommon.NewNotFoundError("menu not found")
	}
	return menu, err
}

func (r postgreSqlMenuRepositoryImpl) FindByName(ctx context.Context, name string) (menu menuDomain.Menu, err error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, name, price, category_name, image_filename, created_at, updated_at, deleted_at FROM menus WHERE name = $1 LIMIT 1;",
		name)

	err = row.Scan(
		&menu.ID,
		&menu.Name,
		&menu.Price,
		&menu.CategoryName,
		&menu.ImageFilename,
		&menu.CreatedAt,
		&menu.UpdatedAt,
		&menu.DeletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return menu, errorCommon.NewNotFoundError("menu not found")
	}
	return menu, err
}

func (r postgreSqlMenuRepositoryImpl) Update(ctx context.Context, menu menuDomain.Menu) (id string, err error) {
	//TODO implement me
	panic("implement me")
}

func (r postgreSqlMenuRepositoryImpl) Delete(ctx context.Context, menuID string) (rid string, err error) {
	//TODO implement me
	panic("implement me")
}
