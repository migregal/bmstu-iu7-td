package user

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"markup2/markupapi/core/interactors"
	"markup2/markupapi/core/ports/repositories"
)

type User struct {
	ID           uint64
	Login        string
	PasswordHash string
}

type UserInfo struct {
	Login    string
	Password string
}

type Interactor struct {
	repo repositories.UserRepo
}

func New(repo repositories.UserRepo) Interactor {
	return Interactor{repo: repo}
}

func (i *Interactor) Register(user UserInfo) (uint64, error) {
	u := repositories.User{
		Login:        user.Login,
		PasswordHash: i.hashPassword(user.Password),
	}

	id, err := i.repo.Create(u)
	if err != nil {
		if errors.Is(err, repositories.ErrExists) {
			err = interactors.ErrExists
		}

		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (i *Interactor) CheckAuth(user User, password string) (bool, error) {
	return i.hashPassword(password) == user.PasswordHash, nil
}

func (i *Interactor) Get(login string) (User, error) {
	user, err := i.repo.Get(login)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			err = interactors.ErrNotFound
		}

		return User{}, fmt.Errorf("failed to get user info: %w", err)
	}

	u := User{ID: user.ID, Login: user.PasswordHash, PasswordHash: user.PasswordHash}
	return u, nil
}

func (i *Interactor) hashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", sum)
}
