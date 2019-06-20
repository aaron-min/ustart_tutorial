package record

import "github.com/aaron-min/ustart_tutorial/record/storage"

// Config determines the runtime behavior of the Elastic-backed record server
type Config struct {
	useDummy      bool
	StorageConfig *storage.Config
}
