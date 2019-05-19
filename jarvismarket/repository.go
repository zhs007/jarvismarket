package jarvismarket

import (
	"bytes"
	"html/template"
	"os/exec"
	"path"

	jarvisbase "github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// RepositoryScriptParam - the parameter for repository script
type RepositoryScriptParam struct {
	RepositoryRootPath string
	RepositoryURL      string
	RepositoryName     string
	RepositoryPath     string
}

func initRepository(cfg *Config, cfgRepo *RepositoryConfig) error {
	gitpath := path.Join(cfg.RepositoryRootPath, cfgRepo.Name, ".git")
	if IsExistsDir(gitpath) {
		runRepositoryScript(cfg, cfgRepo, cfg.OnUpdRepository)
	} else {
		runRepositoryScript(cfg, cfgRepo, cfg.OnInitRepository)
	}

	return nil
}

func runRepositoryScript(cfg *Config, cfgRepo *RepositoryConfig, script string) error {
	tpl, err := template.New("runRepositoryScript").Parse(script)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	params := RepositoryScriptParam{
		RepositoryRootPath: cfg.RepositoryRootPath,
		RepositoryURL:      cfgRepo.URL,
		RepositoryName:     cfgRepo.Name,
		RepositoryPath:     path.Join(cfg.RepositoryRootPath, cfgRepo.Name),
	}

	tpl.Execute(&b, params)

	cmd := exec.Command("/bin/sh", "-c", string(b.Bytes()))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	jarvisbase.Info("runRepositoryScript",
		zap.String("script", string(b.Bytes())),
		zap.String("output", string(output)))

	return nil
}

// InitAllRepositories -
func InitAllRepositories(cfg *Config) error {
	for _, v := range cfg.Repositories {
		err := initRepository(cfg, &v)
		if err != nil {
			jarvisbase.Warn("InitAllRepositories",
				zap.Error(err))

			return err
		}
	}

	return nil
}
