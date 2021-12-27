package worker

import (
	"fmt"
	"log"
	"os"
)

func saveToLocal(savePath string) {
	if exists(savePath) {
		log.Println("文件存在", savePath)
		err := os.Remove(savePath)
		if err != nil {
			fmt.Println(err)
		}
	}

	_, err := os.Create(savePath)
	if err != nil {
		fmt.Println(err)
	}

	//_, err = io.Copy(f, res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
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
