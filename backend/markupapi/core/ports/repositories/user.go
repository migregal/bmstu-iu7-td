package repositories

type UserConfig struct {
	Host      string
	Port      int
	User      string
	Passsword string
	Name      string
}

type User struct {
	ID       uint64
	Login    string
	Password string
}

type UserRepo interface {
	Create(User) error
	Get(id uint64) (User, error)
}
