package client

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
	"os"
	"time"
	"winget-sync/constants"
)

type WebCenter struct {
	Client *ftp.ServerConn
}

func NewWebCenter() *WebCenter {
	c, err := ftp.Dial(constants.FtpAddress, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("lyj", "123456")
	if err != nil {
		log.Fatal(err)
	}

	return &WebCenter{Client: c}
}

func (w *WebCenter) name() {

}

func (w *WebCenter) Upload(name string) error {
	f, err := os.Open(constants.DownloadDir + name)
	if err != nil {
		return err
	}

	err = w.Client.StorFrom(constants.WebDavBasePath+name, f, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func (w *WebCenter) List() (files []*File) {
	list, err := w.Client.List(constants.WebDavBasePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range list {
		files = append(files, &File{
			Name: entry.Name,
			Size: entry.Size,
		})
	}

	return
}

func (w WebCenter) Delete(name string) {
	w.Client.Delete(constants.WebDavBasePath + name)
}

type File struct {
	Name string
	Size uint64
}
