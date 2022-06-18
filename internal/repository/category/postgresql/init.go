package postgresql

import (
	"context"
	"database/sql"
	"errors"
	errorCommon "github.com/aryahmph/restaurant/common/error"
	categoryDomain "github.com/aryahmph/restaurant/internal/domain/category"
)

type postgreSqlCategoryRepositoryImpl struct {
	db *sql.DB
}

func NewPostgreSqlCategoryRepositoryImpl(db *sql.DB) postgreSqlCategoryRepositoryImpl {
	return postgreSqlCategoryRepositoryImpl{db: db}
}

func (r postgreSqlCategoryRepositoryImpl) Insert(ctx context.Context, category categoryDomain.Category) (id string, err error) {
	row := r.db.QueryRowContext(ctx, "INSERT INTO categories(name) VALUES ($1) RETURNING id;", category.Name)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("category not found")
	}
	return id, err
}

func (r postgreSqlCategoryRepositoryImpl) FindAll(ctx context.Context) (categories []categoryDomain.Category, err error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, created_at, updated_at FROM categories;")
	if err != nil {
		return categories, err
	}
	defer rows.Close()
	for rows.Next() {
		var category categoryDomain.Category
		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return categories, err
	}
	return categories, nil
}

func (r postgreSqlCategoryRepositoryImpl) FindByID(ctx context.Context, id string) (category categoryDomain.Category, err error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, name, created_at, updated_at FROM categories WHERE id = $1 LIMIT 1;", id)

	err = row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return category, errorCommon.NewNotFoundError("category not found")
	}
	return category, err
}

func (r postgreSqlCategoryRepositoryImpl) FindByName(ctx context.Context, name string) (category categoryDomain.Category, err error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, name, created_at, updated_at FROM categories WHERE name = $1 LIMIT 1;", name)

	err = row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return category, errorCommon.NewNotFoundError("category not found")
	}
	return category, err
}

func (r postgreSqlCategoryRepositoryImpl) Update(ctx context.Context, category categoryDomain.Category) (id string, err error) {
	row := r.db.QueryRowContext(
		ctx, "UPDATE categories SET name=$1, updated_at= NOW() WHERE id=$2 RETURNING id;",
		category.Name, category.ID)

	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("category not found")
	}
	return id, err
}

func (r postgreSqlCategoryRepositoryImpl) Delete(ctx context.Context, id string) (rid string, err error) {
	row := r.db.QueryRowContext(ctx, "DELETE FROM categories WHERE id=$1 RETURNING id;", id)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return rid, errorCommon.NewNotFoundError("category not found")
	}
	return rid, err
}
