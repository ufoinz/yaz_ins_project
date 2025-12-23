package persistence

import (
	"todo-app/internal/domain/event"

	"gorm.io/gorm"
)

// Event implementation.GORM/Postgres-based repository
type PostgresEventRepo struct{ DB *gorm.DB }

// Creates a new event repository using the transferred GORM connection
func NewPostgresEventRepo(db *gorm.DB) event.Repository {
	return &PostgresEventRepo{DB: db}
}

// saves the new event e to the events table
func (r *PostgresEventRepo) Create(e *event.Event) error {
	return r.DB.Create(e).Error
}

// returns all events from the events table
func (r *PostgresEventRepo) GetAll() ([]event.Event, error) {
	var list []event.Event
	if err := r.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// searches for an event by its ID in the table
func (r *PostgresEventRepo) GetByID(id int64) (*event.Event, error) {
	var ev event.Event
	if err := r.DB.First(&ev, id).Error; err != nil {
		return nil, err
	}
	return &ev, nil
}

// saves changes to object e in the database
func (r *PostgresEventRepo) Update(e *event.Event) error {
	return r.DB.Save(e).Error
}

// deletes an event by its ID
func (r *PostgresEventRepo) Delete(id int64) error {
	return r.DB.Delete(&event.Event{}, id).Error
}
