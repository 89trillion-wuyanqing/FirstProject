package utils

import (
	"FirstProject/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

/**
根据json文件路径读取士兵信息
*/
func GetJson(filePath string) (map[string]model.Army, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var army map[string]model.Army
	e := json.Unmarshal(byteValue, &army)
	return army, e

}

/**
根据士兵信息创建json文件并保存
*/
func CreateJsonFile(army map[string]model.Army) error {
	// 创建文件
	filePtr, err := os.Create("config/info.json")
	if err != nil {
		log.Println("文件创建失败", err.Error())

		return err
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(army)
	if err != nil {
		log.Println("编码错误", err.Error())
		//fmt.Println("编码错误", err.Error())
	} else {
		log.Println("编码成功")
		//fmt.Println("编码成功")
	}

	return err

}

/**
获取保存json文件士兵信息
*/
func GetNewJson() (map[string]model.Army, error) {
	jsonFile, err := os.Open("config/info.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var army map[string]model.Army
	e := json.Unmarshal(byteValue, &army)
	return army, e

}
