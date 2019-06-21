package storage

import (
	"context"

	"github.com/aaron-min/ustart_tutorial/record/recordpb"
)

// Storage is a database-agnostic interface for persisting customer data
type Storage interface {
	New(context.Context, string, string, string, string) error
	History(context.Context, string) (recordpb.Record, error)
	Pay(context.Context, string, float) (float, error)
	// rest of the functions
	ErrUserDoesNotExist() error
	ErrTooManyResults() error
}
