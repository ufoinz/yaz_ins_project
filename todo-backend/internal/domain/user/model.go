package user

// User entity
type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `gorm:"not null" json:"-"`
	Name     string `json:"name" bindeing:"required"`
}

// RegisterRequest — DTO for requesting registration of a new user
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// LoginRequest — DTO for authorization request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
