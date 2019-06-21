package record

import (
	"github.com/aaron-min/ustart_tutorial/record/storage"
)

//Record is an implementation of the record service as defined in service.proto
type Record struct {
	strg          storage.Storage
	defaultAvatar string
	defaultBanner string
}

// New returns a new Eclient record service
func New(cfg *Config) (*Record, error) {
	// if cfg.useDummy

	return nil, nil
}
