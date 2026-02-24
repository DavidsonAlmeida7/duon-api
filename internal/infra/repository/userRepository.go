package repository

import (
	"context"
	"duon-api/internal/core/domain"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id uint) (*domain.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, email, created_at
		FROM user
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, name, email, created_at
		FROM "user"
		WHERE id = $1
	`, id)

	var user domain.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, errors.New("not found")
	}

	return &user, nil
}
