package jarvismarket

import "errors"

var (
	// ErrNoServerAddress - no server address
	ErrNoServerAddress = errors.New("no server address")
	// ErrNoHTTPServerAddr - no http server address
	ErrNoHTTPServerAddr = errors.New("no http server address")
	// ErrNoRepositories - no repositories
	ErrNoRepositories = errors.New("no repositories")
	// ErrNoRepositoriyRootPath - no repositoriy root path
	ErrNoRepositoriyRootPath = errors.New("no repositoriy root path")
	// ErrNoOnInitRepository - no oninitrepository
	ErrNoOnInitRepository = errors.New("no oninitrepository")
	// ErrNoOnUpdRepository - no onupdrepository
	ErrNoOnUpdRepository = errors.New("no onupdrepository")
)
