package utils

import (
	"FirstProject/internal/model"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"os"
)

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

func CreateJsonFile(army map[string]model.Army) error {
	// 创建文件
	filePtr, err := os.Create("conf/info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return err
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(army)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}

	return err

	/*b, err := json.Marshal(army)
	if err != nil {
		return err
	}
	//生成json文件
	err = ioutil.WriteFile("conf/test.json", b, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil*/
}

func GetNewJson() (map[string]model.Army, error) {
	jsonFile, err := os.Open("conf/info.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var army map[string]model.Army
	e := json.Unmarshal(byteValue, &army)
	return army, e

}
