package carapi

import (
	"github.com/aaron-min/ustart_tutorial/car"
)

// Config for auth api
type Config struct {
	CarCfg *car.Config
	Port   int //Recomended 5101
}
