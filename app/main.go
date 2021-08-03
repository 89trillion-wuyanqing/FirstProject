package main

import (
	"FirstProject/app/http"
	"FirstProject/internal/utils"
	flag "github.com/spf13/pflag"
	"log"
	"strings"
)

var fileName = flag.StringP("file", "f", "config/config.army.model.json", "Input json file path")

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func main() {
	//初始化配置
	utils.FileInit()
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	//fmt.Println("fileName=", *fileName)
	returnmap, err := utils.GetJson(*fileName)
	if err != nil {
		log.Fatal("ERROR:" + err.Error())
	} else {
		//fmt.Println(returnmap)
		e := utils.CreateJsonFile(returnmap)
		if e != nil {
			log.Fatal(e.Error())
		}
	}
	//初始化服务器
	http.Init()

}
