package car

import (
	"github.com/aaron-min/ustart_tutorial/car"
)

// RESTAPI implements a REST api
// as a wrapper around the car package.
type RESTAPI struct {
	prof *car.Car
}

// New creates a new car api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	prof, err := car.New(cfg.ProfCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		prof: prof,
	}, nil
}
