package containers

import "time"

type APIClasses struct {
	Name      string     `json:"name" validate:"required,min=2,max=255"`
	StartDate *time.Time `json:"start_date" validate:"required"`
	EndDate   *time.Time `json:"end_date" validate:"required"`
	Capacity  *int       `json:"capacity" validate:"required,numeric,min=1,max=50"`
}
