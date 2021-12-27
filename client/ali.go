package client

import (
	"fmt"
	"github.com/studio-b12/gowebdav"
	"log"
	"os"
	"winget-sync/constants"
	"winget-sync/utils"
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

	//f, err := os.Open(constants.DownloadDir + name)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//err = ali.Client.WriteStream(constants.WebDavBasePath+name, f, 0644)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//f.Close()
}

func (ali *Ali) List() (files []*File) {
	fs, err := ali.Client.ReadDir(constants.WebDavBasePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, f := range fs {
		files = append(files, &File{Size: uint64(f.Size()), Name: f.Name()})
	}
	return
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
		alistName = append(alistName, info.Name)
	}

	fs, err := os.ReadDir(constants.DownloadDir)
	if err != nil {
		fmt.Println(err)
		return

	}

	for _, f := range fs {
		llistName = append(llistName, f.Name())
	}
	del, update := utils.CheckDiff(llistName, alistName)

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
	root := constants.AliWebDavAddr
	user := ""
	password := ""
	ali := gowebdav.NewClient(root, user, password)
	_, err := ali.ReadDir("test")
	if err != nil {
		log.Panicln("阿里云盘连接错误", err)
	}
	return &Ali{
		Client: ali,
	}
}
