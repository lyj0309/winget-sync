package service

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
	"winget-sync/client"
	"winget-sync/constants"
	"winget-sync/model"
	"winget-sync/worker"
)

var db *gorm.DB

func init() {
	var err error
	err = downloadDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("下载数据库完成")
	db, err = gorm.Open(sqlite.Open(constants.DbDsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	//tables, err := db.Migrator().GetTables()
	//for _, table := range tables {
	//	fmt.Println(table)
	//}
}

//getLatestAppList
func getLatestAppList() model.LatestApps {
	var mainFests []model.Manifest
	//db.Select("id", "version", "pathpart").Order("id").Order("version").Where("")
	db.Model(&model.Manifest{}).Select("id, min(version) as version, pathpart").Group("id").Limit(10).Find(&mainFests)

	apps := make(model.LatestApps)

	versionMap := make(map[int]string)
	var versionList []model.Versions
	db.Find(&versionList)
	for _, v := range versionList {
		versionMap[v.Rowid] = v.Version
	}

	idMap := make(map[int]string)
	var ids []model.IDs
	db.Find(&ids)
	for _, i := range ids {
		idMap[i.Rowid] = i.ID
	}

	for _, fest := range mainFests {
		apps[idMap[fest.ID]] = model.AppInfo{
			Version:    versionMap[fest.Version],
			PathpartId: fest.Pathpart,
		}
	}

	return apps
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1
	// select time,max(total) as total,name from timeand group by time ;//取记录中total最大的值

}

//获取app下载路劲
func getUrlPath(row int) (fullPath string) {

	var queryPath func(int)
	queryPath = func(r int) {
		var path model.Pathparts
		db.Where("rowid = ?", r).Find(&path)
		//fmt.Println("path", path)
		if path.Pathpart == "" {
			return
		}
		fullPath = "/" + path.Pathpart + fullPath
		queryPath(path.Parent)
	}
	queryPath(row)
	//fmt.Println(fullPath)
	return
}

//checkUpdate 检查更新并删除老版本
func checkUpdate(apps model.LatestApps) (updates []string) {
	//dir, err := os.ReadDir(constants.DownloadDir)
	//if err != nil {
	//	return nil
	//}
	dir := client.WBc.List()

	appVerCheck := make(map[string]bool)
	for id := range apps {
		appVerCheck[id] = false
	}

	for _, entry := range dir {
		fileName := entry.Name
		idIdx := strings.Index(fileName, "-")
		verIdx := strings.LastIndex(fileName, ".")
		if idIdx == -1 || verIdx == -1 {
			continue
		}

		id := fileName[:idIdx]
		ver := fileName[idIdx+1 : verIdx]

		//检查是否最新版本
		if apps[id].Version == ver {
			appVerCheck[id] = true
			log.Println(id, ver, "最新版")
		} else {
			log.Println(entry.Name, "删除")
			client.WBc.Delete(entry.Name)
			//os.Remove(constants.DownloadDir + entry.Name())
		}
		//entry.Name()
	}
	for s, check := range appVerCheck {
		if !check {
			updates = append(updates, s)
		}
	}
	return
}

func Start() {
	appList := getLatestAppList()
	updates := checkUpdate(appList)
	fmt.Println(updates)

	for _, id := range updates {
		//fmt.Println(id)
		appPath := getUrlPath(appList[id].PathpartId)
		fmt.Println(constants.CacheUrl + appPath)
		worker.NewDownloadTask(constants.CacheUrl + appPath)
		//return
	}
	time.Sleep(time.Hour)
}
