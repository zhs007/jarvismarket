package jarvismarket

import (
	jarvisbase "github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// Init -
func Init(config string) error {
	cfg, err := LoadConfig(config)
	if err != nil {
		jarvisbase.Warn("Init",
			zap.Error(err))

		return err
	}

	return InitAllRepositories(cfg)
}
