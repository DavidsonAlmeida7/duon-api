package repository

import (
	"context"
	"duon-api/internal/core/domain"
	"duon-api/internal/core/domain/helper"
	"duon-api/internal/core/ports/repositories"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repositories.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, nome, email, criado_em
		FROM usuarios
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

	helper.Debug("A", users)

	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, nome, email, criado_em
		FROM usuarios
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
