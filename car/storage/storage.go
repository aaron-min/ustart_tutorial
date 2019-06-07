package storage

import (
	"context"

	"github.com/aaron-min/ustart_tutorial/car/carpb"
)

// Storage is a database-agnostic interface for persisting customer data
type Storage interface {
	Register(context.Context, string, string, string, string) error
	Lookup(context.Context, string) (carpb.Car, error)
	Search(context.Context, []string, bool, map[string][]string, string) ([]string, error)
	// rest of the functions
	ErrUserDoesNotExist() error
	ErrTooManyResults() error
}
