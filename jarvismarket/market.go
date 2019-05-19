package jarvismarket

import (
	"context"

	jarvisbase "github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// Market -
type Market struct {
	cfg *Config
	mgr *RepositoriesMgr
}

// NewMarket -
func NewMarket(filename string) (*Market, error) {
	cfg, err := LoadConfig(filename)
	if err != nil {
		jarvisbase.Warn("Init",
			zap.Error(err))

		return nil, err
	}

	err = InitAllRepositories(cfg)
	if err != nil {
		return nil, err
	}

	return &Market{
		cfg: cfg,
		mgr: &RepositoriesMgr{},
	}, nil
}

// Start -
func (market *Market) Start(ctx context.Context) {

}
