package domain

import "time"

type CompanyConfig struct {
	ID        string
	ApiKey    string // very much work in progress
	CreatedAt time.Time
	UpdatedAt time.Time
}
