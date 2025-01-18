package domain

import "time"

// UserUsage represents the usage of a user in the system in terms of number of
// requests.
type UserUsage struct {
	LastReset    *time.Time
	RequestCount int
	UserID       string
}
