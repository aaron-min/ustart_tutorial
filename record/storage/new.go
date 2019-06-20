package storage

import (
	elasticstore "github.com/aaron-min/ustart_tutorial/record/storage/sql"
)

// NewSQL determines the runtime behavior of the SQL-backed record server
func NewSQL()(config *Config) (Storage, error) {
	strg, err := &Config{SQLConfig: sqlstore.NewConfig()}
	return strg, err
}