package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"winget-sync/model"
)

func main() {
	//dbByte := downloadDb()
	//err := os.WriteFile("index.db", dbByte, 0644)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	db, err := gorm.Open(sqlite.Open("index.db?cache=shared&mode=memory"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return
	}
	for _, table := range tables {
		fmt.Println(table)
	}
	var mf []model.Manifest
	db.Find(&mf)
	for _, manifest := range mf {
		fmt.Println(manifest.ID)
	}
	//匹配 1，名字. 2.tag 3.monikers

}

func downloadDb() []byte {
	resp, err := http.Get("https://winget.azureedge.net/cache/source.msix")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Fatal(err)
	}
	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		fmt.Println("Reading file:", zipFile.Name)
		if zipFile.Name == "Public/index.db" {
			unzippedFileBytes, err := readZipFile(zipFile)
			if err != nil {
				log.Println(err)
				return nil
			}
			return unzippedFileBytes
		}

		//_ = unzippedFileBytes // this is unzipped file bytes
	}
	return nil
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
