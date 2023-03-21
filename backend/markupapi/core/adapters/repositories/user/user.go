package user

import (
	"fmt"
	"markup2/markupapi/core/ports/repositories"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID           uint64 `gorm:"primaryKey"`
	Login        string `gorm:"index"`
	PasswordHash string
}

type Repository struct {
	db *gorm.DB
}

func New(cfg repositories.UserConfig) (*Repository, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Name, cfg.User, cfg.Passsword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Create(user repositories.User) (uint64, error) {
	u := User{Login: user.Login, PasswordHash: user.PasswordHash}

	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return u.ID, nil
}

func (r *Repository) Get(login string) (repositories.User, error) {
	return repositories.User{}, nil
}
