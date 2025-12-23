package event

// The abstraction interface over the event store
// Allows you to work with events
type Repository interface {
	Create(e *Event) error
	GetAll() ([]Event, error)
	GetByID(id int64) (*Event, error)
	Update(e *Event) error
	Delete(id int64) error
}
