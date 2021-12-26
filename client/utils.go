package client

import "fmt"

//CheckDiff 检查两个文件夹的不同,返回 del-要删除的 , update-要上传的
func CheckDiff(src []string, tg []string) ([]string, []string) {
	for i1 := 0; i1 < len(src); i1++ {
		for i2, s2 := range tg {
			if src[i1] == s2 {
				//fmt.Println(s, src)
				src = append(src[:i1], src[i1+1:]...)
				tg = append(tg[:i2], tg[i2+1:]...)
				i1--
				//fmt.Println(tg)
				fmt.Println(src)
				break
			}
		}
	}
	return tg, src
}
