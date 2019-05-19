package jarvismarket

import (
	"path"
	"path/filepath"
	"sync"

	jarvisbase "github.com/zhs007/jarviscore/base"
	jarvismarketpb "github.com/zhs007/jarvismarket/proto"
	"go.uber.org/zap"
)

// appList -
type appList struct {
	lst []*jarvismarketpb.AppInfo
}

// RepositoriesMgr -
type RepositoriesMgr struct {
	sync.Mutex

	lst []*jarvismarketpb.AppInfo
}

// FindAppWithKeyword -
func (mgr *RepositoriesMgr) FindAppWithKeyword(key string) []*jarvismarketpb.AppInfo {
	mgr.Lock()
	defer mgr.Unlock()

	var lst []*jarvismarketpb.AppInfo

	for _, v := range mgr.lst {
		if HasKeyword(v, key) {
			lst = append(lst, v)
		}
	}

	return lst
}

// Reload -
func (mgr *RepositoriesMgr) Reload(cfg *Config) error {
	mgr.Lock()
	defer mgr.Unlock()

	lst, err := filepath.Glob(path.Join(cfg.RepositoryRootPath, "**/*.yaml"))
	if err != nil {
		jarvisbase.Warn("RepositoriesMgr.Reload",
			zap.Error(err))

		return err
	}

	for _, v := range lst {
		appinfo, err := LoadAppInfo(v)
		if err != nil {
			jarvisbase.Warn("RepositoriesMgr.Reload",
				zap.Error(err))
		} else {
			msg := appinfo.ToProto()

			err = mgr.updAppInfo(msg)
			if err != nil {
				jarvisbase.Warn("RepositoriesMgr.Reload:updAppInfo",
					zap.Error(err))
			}
		}
	}

	return nil
}

// updAppInfo -
func (mgr *RepositoriesMgr) updAppInfo(appinfo *jarvismarketpb.AppInfo) error {

	for i, v := range mgr.lst {
		if v.Name == appinfo.Name && v.Author == appinfo.Author {
			lst := append(mgr.lst[0:i], appinfo)
			if i < len(mgr.lst) {
				mgr.lst = append(lst, mgr.lst[i+1:]...)
			}

			return nil
		}
	}

	mgr.lst = append(mgr.lst, appinfo)

	return nil
}
