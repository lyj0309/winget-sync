package worker

import (
	"fmt"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
	"winget-sync/client"
	"winget-sync/constants"
	"winget-sync/model"
)

var p pool.Pool
var ali *client.Ali
var wb *client.WebCenter
var count int

func init() {
	p = pool.NewLimited(10)
	//ali = client.NewAli()

	wb = client.NewWebCenter()
}

func downloadExe(url string, savePath string) {
	fmt.Println(url, savePath)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	if exists(savePath) {
		log.Println("文件存在", savePath)
		err = os.Remove(savePath)
		if err != nil {
			fmt.Println(err)
		}
	}

	f, err := os.Create(savePath)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		f.Close()
		res.Body.Close()
	}()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//exec.Command("mv", savePath, "/root/ftp/wg/").Run()

}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

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

	//加速

	return

}

func NewDownloadTask(url string) {
	//加入下载队列
	p.Queue(func(wu pool.WorkUnit) (data interface{}, err error) {
		info, err := getAppInfo(url)
		if err != nil {
			return
		}
		ddUrl, suffix := info.GetAppDownloadInfo()

		//return
		pkgName := info.PackageIdentifier + "-" + info.PackageVersion + "." + suffix
		downloadExe(ddUrl, constants.DownloadDir+pkgName)
		log.Println(pkgName, "下载完成")

		log.Println(pkgName + "上传完成")

		count++
		fmt.Println(count)
		return
	})
}
