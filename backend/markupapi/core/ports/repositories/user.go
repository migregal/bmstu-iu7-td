package repositories

type UserConfig struct {
	Host      string
	Port      int
	User      string
	Passsword string
	Name      string
}

type User struct {
	ID           uint64
	Login        string
	PasswordHash string
}

type UserRepo interface {
	Create(User) (uint64, error)
	Get(login string) (User, error)
}
