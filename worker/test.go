package worker

import (
	"fmt"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var total = 0
var poolCount = 0
var p = pool.NewLimited(10)

func GetTreeSize(path string) (int, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	//判断文件夹是不是版本号
	if len(entries) > 0 {
		if entries[0].IsDir() {
			if isVer(entries[0].Name()) {
				//找到最新版本
				for i := len(entries); i > 0; i-- {
					if isVer(entries[i-1].Name()) {
						total++
						//fmt.Println(filepath.Join(path, entries[i-1].Name()))
						findDownload(filepath.Join(path, entries[i-1].Name()))
						return total, nil
					}
				}
				return total, nil
			}
		}
	}

	for _, entry := range entries {
		if entry.IsDir() {
			_, err := GetTreeSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return 0, err
			}
		}
	}
	return total, nil
}

func findDownload(path string) {
	//fmt.Println(path)
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, e := range entries {
		f, err := os.ReadFile(filepath.Join(path, e.Name()))
		if err != nil {
			return
		}
		//fmt.Println(string(f))
		t := T{}
		_ = yaml.Unmarshal(f, &t)
		if len(t.Installers) > 0 {
			if len(t.Installers[0].InstallerUrl) > 0 {
				downloadFile(t.Installers[0].InstallerUrl, path)
				//fmt.Println(t.Installers[0].InstallerUrl, path)
			}
		}
	}
}
func downloadFile(url string, savePath string) {
	i := strings.LastIndex(url, "/")
	url = strings.ReplaceAll(url, "github.com", "download.fastgit.org")
	if i != -1 {
		savePath = savePath + "/" + url[i+1:]
	} else {
		return
	}

	p.Queue(func(wu pool.WorkUnit) (interface{}, error) {
		fmt.Println(url, savePath)
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}
		os.Remove(savePath)
		f, err := os.Create(savePath)
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}
		io.Copy(f, res.Body)

		exec.Command("mv", savePath, "/root/ftp/wg/").Run()
		if wu.IsCancelled() {
			// return values not used
			return nil, nil
		}
		return nil, nil
	})
}

//isVer 判断版本号
func isVer(s string) bool {
	return regexp.MustCompile(`\d*\.\d*`).MatchString(s)
}

type T struct {
	InstallerType string `yaml:"InstallerType"`
	//License    string    `yaml:"License"`
	Installers []insInfo `yaml:"Installers"`
}

type insInfo struct {
	InstallerUrl string `yaml:"InstallerUrl"`
}

func NewDownloadTask(url string) {

}
