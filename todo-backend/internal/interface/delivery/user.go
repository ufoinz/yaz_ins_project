package delivery

import (
	"errors"
	"net/http"
	"todo-app/internal/domain/user"
	"todo-app/internal/infrastructure/security"

	"github.com/gin-gonic/gin"
)

// UserHandler stores a service for working with users and a secret for JWT
type UserHandler struct {
	svc       user.Service
	jwtSecret string
}

// routes:
func NewUserHandler(r *gin.RouterGroup, svc user.Service, jwtSecret string) {
	h := &UserHandler{svc: svc, jwtSecret: jwtSecret}
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}

// Register handles the POST /users/register
func (h *UserHandler) Register(c *gin.Context) {
	var req user.RegisterRequest

	// binding and validation of fields
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// user registration via the service
	u, err := h.svc.Register(req)
	if err != nil {
		if errors.Is(err, user.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, u)
}

// Login handles the POST /users/login
func (h *UserHandler) Login(c *gin.Context) {
	var req user.LoginRequest

	// binding and validation of fields
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// user authentication
	u, err := h.svc.Authenticate(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, user.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// generation of a JWT with a user_id field and a 24-hour validity period
	token, err := security.GenerateToken(u.ID, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
