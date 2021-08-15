package utils

import (
	"io/ioutil"
	"strings"
)

//获取文件名字
func GetFileName(dir string, username string) (filename string, exist bool) {
	files,err:=ioutil.ReadDir(dir)
	if err != nil {
		return "",false
	}
	for _,f:=range files {
		filename = f.Name()
		if strings.Contains(filename, username) {
			return filename,true
		}
	}
	return "",false
}
