package incident

import (
	"time"
	"github.com/google/uuid"
)

// Incident is a basic 
type Incident struct {
	Title string
	Description string
	CreatedAt time.Time
	ID string
}

// Persist is a stub, in reality it would save the incident to a database.
func Persist(inc *Incident) error {
	inc.CreatedAt = time.Now()
	inc.ID = uuid.New().String()

	return nil
}