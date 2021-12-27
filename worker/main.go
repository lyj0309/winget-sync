package worker

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"gopkg.in/go-playground/pool.v3"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
	"winget-sync/client"
	"winget-sync/constants"
	"winget-sync/model"
)

var p pool.Pool

var count int

func init() {
	p = pool.NewLimited(10)
	//ali = client.NewAli()
}

//download 下载文件
func download(url string, path string) error {
	// create client
	c := grab.NewClient()
	req, _ := grab.NewRequest(path, url)
	// start download
	//fmt.Printf("下载 %v中...\n", req.URL())
	resp := c.Do(req)

	//fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("%s , %.2f m/s(%.2f%%) \n",
				req.URL(),
				resp.BytesPerSecond()/1024/1024,
				100*resp.Progress(),
			)

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "下载失败: %v\n", err)
		return err
	}

	fmt.Printf("下载完成 ,saved to %v \n", resp.Filename)
	return nil
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
		fileName := info.PackageIdentifier + "-" + info.PackageVersion + "." + suffix
		//constants.DownloadDir+pkgName
		err = download(ddUrl, constants.DownloadDir+fileName)
		if err != nil {
			return
		}

		client.WBc.Upload(fileName)
		log.Println(fileName + "上传ftp完成")
		//client.Alic.Upload(fileName)
		//log.Println(fileName + "上传ali完成")

		count++
		fmt.Println(count)
		return
	})
}
