package client

import (
	"fmt"
	"github.com/studio-b12/gowebdav"
	"log"
	"os"
	"winget-sync/constants"
)

type Ali struct {
	Client *gowebdav.Client
}

func (ali *Ali) Upload(name string) {

	b, err := os.ReadFile(constants.DownloadDir + name)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ali.Client.Write(constants.WebDavBasePath+name, b, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (ali *Ali) List() []os.FileInfo {
	files, err := ali.Client.ReadDir(constants.WebDavBasePath)
	if err != nil {
		fmt.Println(err)
	}
	return files
}

func (ali *Ali) Delete(name string) {
	err := ali.Client.Remove(constants.WebDavBasePath + name)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (ali Ali) Sync() {

	var alistName []string
	var llistName []string

	alist := ali.List()
	for _, info := range alist {
		alistName = append(alistName, info.Name())
	}

	fs, err := os.ReadDir(constants.DownloadDir)
	if err != nil {
		fmt.Println(err)
		return

	}

	for _, f := range fs {
		llistName = append(llistName, f.Name())
	}
	del, update := CheckDiff(llistName, alistName)

	for _, s := range del {
		log.Println("删除", s)
		ali.Delete(s)
	}

	for _, s := range update {
		log.Println("上传", s)
		ali.Upload(s)
	}

}

func NewAli() *Ali {
	root := "http://localhost:8080"
	user := ""
	password := ""
	return &Ali{
		Client: gowebdav.NewClient(root, user, password),
	}
}