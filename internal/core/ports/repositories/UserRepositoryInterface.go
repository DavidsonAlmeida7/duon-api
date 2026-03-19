package repositories

import (
	"context"
	"duon-api/internal/core/domain"
)

// UserRepository define como o Core espera que os dados sejam salvos ou buscados.
// Não importa se é Postgres, MongoDB ou um arquivo JSON.
// Output Port: O que o UseCase exige do banco de dados
type UserRepositoryInterface interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	//Save(user domain.User) error
	//ExistsByEmail(email string) (bool, error)
}
