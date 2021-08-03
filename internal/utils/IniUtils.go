package utils

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	//配置信息
	iniFile *ini.File
)

//初始化文件加载信息
func FileInit() {
	file, e := ini.Load("././config/app.ini")
	if e != nil {
		log.Fatal("Fail to load config/app.ini" + e.Error())
	}
	iniFile = file
}

//获取配置中section域
func GetSection(sectionName string) *ini.Section {
	section, e := iniFile.GetSection(sectionName)
	if e != nil {
		log.Fatal("未找到对应的配置信息:" + sectionName + e.Error())

	}
	return section
}

/**
获取对应section下的对应key的字符串值
*/
func GetVal(sectionName string, key string) string {
	var temp_val string
	section := GetSection(sectionName)
	if nil != section {
		temp_val = section.Key(key).Value()
	}
	return temp_val
}
