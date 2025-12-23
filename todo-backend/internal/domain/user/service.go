package user

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Error with incorrect authorization
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailExists        = errors.New("email already exists")
)

// User Business Logic Interface (Registration)
type Service interface {
	Register(req RegisterRequest) (User, error)
	Authenticate(email, password string) (User, error)
}

type service struct {
	repo Repository
}

// User Service Constructor
func NewService(r Repository) Service {
	return &service{repo: r}
}

// hash the password and save the new user
func (s *service) Register(req RegisterRequest) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	u := User{
		Email:    req.Email,
		Password: string(hash),
		Name:     req.Name,
	}

	if err := s.repo.Insert(&u); err != nil {
		if strings.Contains(err.Error(), "uni_users_email") || strings.Contains(err.Error(), "duplicate") {
			return User{}, ErrEmailExists
		}
		return User{}, err
	}

	return u, nil
}

// check the email and password, comparing it with the hash
func (s *service) Authenticate(email, password string) (User, error) {
	u, err := s.repo.GetByEmail(email)
	if err != nil {
		return User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return User{}, ErrInvalidCredentials
	}
	return u, nil
}
