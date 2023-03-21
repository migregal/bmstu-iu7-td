package user

import "markup2/markupapi/core/ports/repositories"

type Interactor struct {
	repo repositories.UserRepo
}

func New(repo repositories.UserRepo) Interactor {
	return Interactor{repo: repo}
}

func (i *Interactor) Register() {

}

func (i *Interactor) Get() {

}
