package main

import (
	"github.com/robfig/cron"
	"winget-sync/service"
)

func main() {
	c := cron.New()
	c.AddFunc("50 1 * * *", func() {
		service.Start()
	})

	go c.Start()
	defer c.Stop()

	select {}
	//匹配 1，名字. 2.tag 3.monikers

}
