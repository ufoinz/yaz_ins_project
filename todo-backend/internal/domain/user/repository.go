package user

// Repository — abstraction interface over user storage
// Allows you to work with data without linking to a database
type Repository interface {
	Insert(u *User) error
	GetByEmail(email string) (User, error)
	GetByID(id int64) (User, error)
}
