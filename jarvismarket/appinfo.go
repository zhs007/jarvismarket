package jarvismarket

import (
	"io/ioutil"
	"os"

	jarvismarketpb "github.com/zhs007/jarvismarket/proto"
	"gopkg.in/yaml.v2"
)

// AppInfo -
type AppInfo struct {
	Author      string
	AuthorEmail string
	Name        string
	Description string
	Keywords    []string
	Version     string
	Type        string
	Docker      bool
}

// LoadAppInfo - load appinfo
func LoadAppInfo(filename string) (*AppInfo, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	appinfo := &AppInfo{}

	err = yaml.Unmarshal(fd, appinfo)
	if err != nil {
		return nil, err
	}

	return appinfo, nil
}

// ToProto -
func (appinfo *AppInfo) ToProto() *jarvismarketpb.AppInfo {
	msg := &jarvismarketpb.AppInfo{
		Author:      appinfo.Author,
		AuthorEmail: appinfo.AuthorEmail,
		Name:        appinfo.Name,
		Description: appinfo.Description,
		InDocker:    appinfo.Docker,
		Version:     appinfo.Version,
		Type:        jarvismarketpb.AppType_APPTYPE_NORMAL,
	}

	for _, v := range appinfo.Keywords {
		msg.Keywords = append(msg.Keywords, v)
	}

	if appinfo.Type == "service" {
		msg.Type = jarvismarketpb.AppType_APPTYPE_SERVICE
	}

	return msg
}
