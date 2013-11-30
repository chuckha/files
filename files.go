package files

import (
	"os"
)

// MustGetOrCreateFile gets an *os.File by opening or creating.
// Used for things like creating a logging file on a new server.
func MustGetOrCreateFile(filename string) *os.File {
	var f *os.File
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if os.IsNotExist(err) {
		f, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
	}
	return f
}
