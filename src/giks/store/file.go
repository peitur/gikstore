package store

import "os"

// Exists reports whether the named file or directory exists.
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
