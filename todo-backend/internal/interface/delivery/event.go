package delivery

import (
	"net/http"
	"strconv"
	"time"
	"todo-app/internal/domain/event"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateEvent handles POST /events
func (app *Application) CreateEvent(c *gin.Context) {
	// expected fields of the JSON request
	var payload struct {
		Name    string    `json:"name" binding:"required,min=1"`
		Content string    `json:"content"`
		Time    time.Time `json:"time" binding:"required"`
	}

	// JSON binding and validation
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// getting the user ID from the JWT middleware
	userID, _ := c.Get("user_id")
	ownerID := userID.(int64)

	// creating the event domain object
	ev := event.Event{
		OwnerId: ownerID,
		Name:    payload.Name,
		Content: payload.Content,
		Time:    payload.Time,
	}

	// saving the event
	if err := app.Events.Create(&ev); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create event"})
		return
	}
	c.JSON(http.StatusCreated, ev)
}

// getEvents handles GET /events
func (app *Application) GetEvents(c *gin.Context) {
	all, err := app.Events.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch events"})
		return
	}

	userID, _ := c.Get("user_id")
	ownerID := userID.(int64)

	// we collect only our own events
	mine := make([]event.Event, 0, len(all))
	for _, e := range all {
		if e.OwnerId == ownerID {
			mine = append(mine, e)
		}
	}
	c.JSON(http.StatusOK, mine)
}

// GetEvent handles GET /events/:id
func (app *Application) GetEvent(c *gin.Context) {
	// parse the id parameter
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	// looking for an event
	ev, err := app.Events.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch event"})
		}
		return
	}

	// checking that the event belongs to the user
	userID, _ := c.Get("user_id")
	if ev.OwnerId != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	c.JSON(http.StatusOK, ev)
}

// UpdateEvent handles PUT /events/:id
func (app *Application) UpdateEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	// loading an existing event
	existing, err := app.Events.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch event"})
		}
		return
	}

	// checking ownership
	userID, _ := c.Get("user_id")
	if existing.OwnerId != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	// new payload for updating
	var payload struct {
		Name    string    `json:"name" binding:"required,min=3"`
		Content string    `json:"content"`
		Time    time.Time `json:"time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// applying the changes
	existing.Name = payload.Name
	existing.Content = payload.Content
	existing.Time = payload.Time

	// saving the updated event
	if err := app.Events.Update(existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, existing)
}

// DeleteEvent handles DELETE /events/:id
func (app *Application) DeleteEvent(c *gin.Context) {
	// parse the ID from the URL
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	// searching for an event by ID
	ev, err := app.Events.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch event"})
		}
		return
	}

	// checking the deletion rights
	userID, _ := c.Get("user_id")
	if ev.OwnerId != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	// deleting the event
	if err := app.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.Status(http.StatusNoContent)
}
