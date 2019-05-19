package jarvismarket

import "errors"

var (
	// ErrNoServerAddress - no server address
	ErrNoServerAddress = errors.New("no server address")
	// ErrNoHTTPServerAddr - no http server address
	ErrNoHTTPServerAddr = errors.New("no http server address")
	// ErrNoRepositories - no repositories
	ErrNoRepositories = errors.New("no repositories ")
)
