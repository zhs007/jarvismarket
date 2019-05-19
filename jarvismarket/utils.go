package jarvismarket

import (
	"os"
	"strings"

	"github.com/zhs007/jarvismarket/proto"
)

// IsExistsDir - Check if this folder exists?
func IsExistsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return s.IsDir()
}

// HasKeyword - has keyword
func HasKeyword(appinfo *jarvismarketpb.AppInfo, key string) bool {
	if strings.Contains(appinfo.Name, key) {
		return true
	}
	
	for _, v := range appinfo.Keywords {
		if strings.Contains(v, key) {
			return true
		}
	}

	return false
}