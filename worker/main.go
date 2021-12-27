package worker

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"winget-sync/client"
	"winget-sync/constants"
	"winget-sync/model"
)

var count int

func getAppInfo(url string) (info model.YamlAppInfo, err error) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	err = yaml.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return
	}
	return

}

func NewDownloadTask(url string) {
	//加入下载队列
	p.Queue(func(wu pool.WorkUnit) (data interface{}, err error) {
		info, err := getAppInfo(url)
		if err != nil {
			logrus.Error(err)
			return
		}
		ddUrl, suffix := info.GetAppDownloadInfo()

		//return
		fileName := info.PackageIdentifier + "-" + info.PackageVersion + "." + suffix
		//constants.DownloadDir+pkgName
		err = download(ddUrl, constants.DownloadDir+fileName)
		if err != nil {
			logrus.Error("下载错误", err)
			return
		}
		err = client.NewWebCenter().Upload(fileName)
		if err != nil {
			logrus.Error("上传ftp失败", err)
			return
		}
		//client.WBc.Upload(fileName)
		log.Println("上传ftp完成", fileName)

		err = client.Alic.Upload(fileName)
		if err != nil {
			logrus.Error("上传ali失败", err)
			return
		}
		log.Println("上传ali完成", fileName)

		count++
		log.Println("下载计数", count)
		return
	})
}
