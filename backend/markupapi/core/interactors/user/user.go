package user

import (
	"fmt"
	"markup2/markupapi/core/ports/repositories"
)

type User struct {
	ID           uint64
	Login        string
	PasswordHash string
}
type Interactor struct {
	repo repositories.UserRepo
}

func New(repo repositories.UserRepo) Interactor {
	return Interactor{repo: repo}
}

func (i *Interactor) Register() (uint64, error) {
	user := repositories.User{}

	id, err := i.repo.Create(user)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (i *Interactor) Get(login string) (User, error) {
	user, err := i.repo.Get(login)
	if err != nil {
		return User{}, fmt.Errorf("failed to get user info: %w", err)
	}

	u := User{ID: user.ID, Login: user.PasswordHash, PasswordHash: user.PasswordHash}
	return u, err
}
