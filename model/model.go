package model

import (
	"fmt"
	"strings"
)

type LatestApps map[string]AppInfo

type AppInfo struct {
	Version    string
	PathpartId int
}
type YamlAppInfo struct {
	InstallerType     string `yaml:"InstallerType"`
	PackageVersion    string `yaml:"PackageVersion"`
	PackageIdentifier string `yaml:"PackageIdentifier"`
	//License    string    `yaml:"License"`
	Installers []insInfo `yaml:"Installers"`
}

type insInfo struct {
	Architecture  string `yaml:"Architecture"` //处理器类型
	InstallerUrl  string `yaml:"InstallerUrl"`
	InstallerType string `yaml:"InstallerType"`
}

func (info YamlAppInfo) GetAppDownloadInfo() (string, string) {
	ddUrl := info.Installers[0].InstallerUrl
	//fmt.Println("appInfo", info)
	suffix := "exe"

	swType := func(t string) {
		switch t {
		case "msi":
			suffix = t
		case "msix":
			suffix = t
		}
	}

	if info.InstallerType == "" {
		swType(info.Installers[0].InstallerType)
	} else {
		swType(info.InstallerType)
	}

	if len(info.Installers) > 1 {
		for _, installer := range info.Installers {
			if installer.Architecture == "x64" {
				ddUrl = installer.InstallerUrl
				break
			}
		}
	}

	fmt.Println(info.PackageIdentifier, suffix)

	ddUrl = strings.ReplaceAll(ddUrl, "github.com", "download.fastgit.org")
	return ddUrl, suffix
}
