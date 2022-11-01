package postgres

import (
	"context"
	"database/sql"
	"plato-tech/muly/domain/users"
	appError "plato-tech/muly/utils/app-error"
	"time"
)

type userRepo struct {
	Conn *sql.DB
}

func (u *userRepo) Insert(user *users.User) error {
	query := `
		INSERT INTO users(name, email, password_hash, activated)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version
	`
	args := []interface{}{user.Name, user.Email, user.Password.Hash, user.Activated}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.Conn.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return appError.ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}
