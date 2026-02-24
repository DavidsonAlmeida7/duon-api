package repositories

import "duon-api/internal/core/domain"

// UserRepository define como o Core espera que os dados sejam salvos ou buscados.
// Não importa se é Postgres, MongoDB ou um arquivo JSON.
type UserRepository interface {
	Save(user domain.User) error
	FindByID(id string) (domain.User, error)
	ExistsByEmail(email string) (bool, error)
}
