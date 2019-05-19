package jarvismarket

import "os"

// IsExistsDir - Check if this folder exists?
func IsExistsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return s.IsDir()
}
