package loadpath

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"

)

type JsonStruct struct {
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
	file, err := os.Create("NewConfig.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(v)
	if err!=nil{
		fmt.Println("err:",err)
	}
}
