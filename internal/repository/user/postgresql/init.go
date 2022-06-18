package postgresql

import (
	"context"
	"database/sql"
	"errors"
	errorCommon "github.com/aryahmph/restaurant/common/error"
	userDomain "github.com/aryahmph/restaurant/internal/domain/user"
)

type postgreSqlUserRepositoryImpl struct {
	db *sql.DB
}

func NewPostgreSqlUserRepositoryImpl(db *sql.DB) postgreSqlUserRepositoryImpl {
	return postgreSqlUserRepositoryImpl{db: db}
}

func (r postgreSqlUserRepositoryImpl) Insert(ctx context.Context, user userDomain.User) (id string, err error) {
	row := r.db.QueryRowContext(ctx, "INSERT INTO users(username, password_hash, role) VALUES ($1, $2, $3) RETURNING id;",
		user.Username,
		user.PasswordHash,
		user.GetUserRoleString(),
	)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}

func (r postgreSqlUserRepositoryImpl) FindAll(ctx context.Context) (users []userDomain.User, err error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, username, password_hash, role, created_at, updated_at FROM users;")
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user userDomain.User
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.PasswordHash,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (r postgreSqlUserRepositoryImpl) FindByID(ctx context.Context, id string) (user userDomain.User, err error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, username, password_hash, role, created_at, updated_at FROM users WHERE id = $1 LIMIT 1;",
		id)

	err = row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errorCommon.NewNotFoundError("user not found")
	}
	return user, err
}

func (r postgreSqlUserRepositoryImpl) FindByUsername(ctx context.Context, username string) (user userDomain.User, err error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, username, password_hash, role, created_at, updated_at FROM users WHERE username = $1 LIMIT 1;",
		username)

	err = row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errorCommon.NewNotFoundError("user not found")
	}
	return user, err
}

func (r postgreSqlUserRepositoryImpl) Update(ctx context.Context, user userDomain.User) (id string, err error) {
	row := r.db.QueryRowContext(
		ctx, "UPDATE users SET username=$1, role=$2, updated_at= NOW() WHERE id=$3 RETURNING id;",
		user.Username, user.Role, user.ID)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}

func (r postgreSqlUserRepositoryImpl) Delete(ctx context.Context, id string) (rid string, err error) {
	row := r.db.QueryRowContext(ctx, "DELETE FROM users WHERE id=$1 RETURNING id;", id)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return rid, errorCommon.NewNotFoundError("user not found")
	}
	return rid, err
}
