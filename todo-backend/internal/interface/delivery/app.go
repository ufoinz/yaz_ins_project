package delivery

import (
	nethttp "net/http"
	"todo-app/internal/domain/event"
	"todo-app/internal/domain/user"
	"todo-app/internal/infrastructure/security"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// stores the global configuration and dependencies of the service
type Application struct {
	Port        int
	JWTSecret   string
	Events      event.Repository
	UserService user.Service
}

// configures all HTTP routes and middleware servers
func (app *Application) Routes() nethttp.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(nethttp.StatusOK, gin.H{"message": "pong"})
	})

	v1 := r.Group("/api/v1")

	// routes for working with users
	users := v1.Group("/users")
	NewUserHandler(users, app.UserService, app.JWTSecret)

	// routes for working with events protected by JWT
	events := v1.Group("/events")
	events.Use(security.JWTMiddleware(app.JWTSecret))
	{
		events.POST("/", app.CreateEvent)
		events.GET("/", app.GetEvents)
		events.GET("/:id", app.GetEvent)
		events.PUT("/:id", app.UpdateEvent)
		events.DELETE("/:id", app.DeleteEvent)
	}

	return r
}
