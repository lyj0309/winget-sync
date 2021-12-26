package worker

import (
	"fmt"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {
	t1 := time.Now()
	NewDownloadTask("https://winget.azureedge.net/cache/manifests/t/Tencent/QQMusic/18.41/9f4b-Tencent.QQMusic.yaml")
	elapsed := time.Since(t1)
	fmt.Println("Appelapsed:", elapsed)

}
