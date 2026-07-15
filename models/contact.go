package models

import "time"

type Contact struct {
	Id int64

	CreatedAt time.Time
	UpdatedAt time.Time

	Name    string
	Dob     *time.Time
	Address *string
	Phone   *string
	Email   *string
}
