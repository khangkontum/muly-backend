package authPostgres

import (
	"context"
	"database/sql"
	"errors"
	"plato-tech/muly/domain"
	"plato-tech/muly/utils/appError"
)

type userRepo struct {
	Conn *sql.DB
}

func NewUserRepo(conn *sql.DB) domain.UserRepository {
	return &userRepo{conn}
}

func (ur *userRepo) Insert(ctx context.Context, user *domain.User) (domain.User, error) {
	query := `
		INSERT INTO users(name, email, password_hash, activated)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version
	`
	args := []interface{}{user.Name, user.Email, user.Password.Hash, user.Activated}

	err := ur.Conn.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return *user, appError.ErrDuplicateEmail
		default:
			return *user, err
		}
	}
	return *user, nil
}

func (ur *userRepo) GetByEmail(ctx context.Context, email string) (user domain.User, err error) {
	query := `
 	 	 SELECT id, created_at, name, email, password_hash, activated, version
 	 	 FROM users
 	 	 WHERE email = $1
 	 `
	err = ur.Conn.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Activated,
		&user.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return domain.User{}, appError.ErrRecordNotFound
		default:
			return domain.User{}, err
		}
	}
	return user, nil
}

func (ur *userRepo) Update(ctx context.Context, user *domain.User) (domain.User, error) {
	var returnUser domain.User
	query := `
		UPDATE users
		SET name = $1, email = $2, password_hash = $3, activated = $4, version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING name, email, password_hash, activated, version
	`

	args := []interface{}{
		user.Name,
		user.Email,
		user.Password.Hash,
		user.Activated,
		user.Version,
	}

	err := ur.Conn.QueryRowContext(ctx, query, args...).Scan(
		&returnUser.Name,
		&returnUser.Email,
		&returnUser.Password.Hash,
		&returnUser.Activated,
		&returnUser.Version,
	)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return domain.User{}, appError.ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return domain.User{}, appError.ErrEditConflict
		default:
			return domain.User{}, err
		}
	}
	return returnUser, nil
}
