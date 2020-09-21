package containers

import "time"

type APIBookings struct {
	Email string     `json:"email" validate:"required,email"`
	Date  *time.Time `json:"date" validate:"required"`
}
