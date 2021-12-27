package service

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"winget-sync/constants"
)

func downloadDb() error {
	resp, err := http.Get(constants.WingetAppSource)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}
	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		//fmt.Println("Reading file:", zipFile.Name)
		if zipFile.Name == "Public/index.db" {
			unzippedFileBytes, err := readZipFile(zipFile)
			if err != nil {
				return err
			}
			os.WriteFile("index.db", unzippedFileBytes, 0644)
			return nil
		}

		//_ = unzippedFileBytes // this is unzipped file bytes
	}
	return errors.New("下载数据库错误")
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)

}

func clearCache() {
	os.RemoveAll(constants.DownloadDir)
	os.Mkdir(constants.DownloadDir, 0644)
}
