package utils

import (
	"awesomeProject/gin_demo/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Soldier map[string]model.Soldier

type JsonStruct struct {
}

func init() {
	//获取命令行json文件路径
	filename := Identify()
	jsonStruct := NewJsonStruct()
	//传入数据到Soldier
	jsonStruct.Load(filename,&Soldier)
}


func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

//解析json文件，数据传入数据结构
func (jsd *JsonStruct) Load(filename string,v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err!=nil {
		return
	}
	//将json数据解码到数据结构
	err = json.Unmarshal(data, v)
	if err!=nil {
		return
	}
	//创建新的json文件
	file, err := os.Create("./config/NewConfig.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	//将新的json数据写入新文件
	encoder := json.NewEncoder(file)
	err = encoder.Encode(v)
	if err!=nil{
		fmt.Println("err:",err)
	}
}
