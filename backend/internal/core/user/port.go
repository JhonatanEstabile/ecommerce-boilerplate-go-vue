package user

type Repository interface {
	FindByEmail(email string) (*User, error)
	Create(user User) error
}

type AuthAdapter interface {
	ComparePassword(hashed string, password string) bool
	GenerateToken(userID int64) (string, error)
	HashPassword(password string) (string, error)
}
