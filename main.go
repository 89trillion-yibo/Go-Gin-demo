package main

import (
	"awesomeProject/gin_demo/loadpath"
	"strconv"

	"awesomeProject/gin_demo/pflag"
	"awesomeProject/gin_demo/soldiermethod"
	"github.com/gin-gonic/gin"
	"github.com/tietang/props/ini"
)


type Vmap map[string]soldiermethod.Soldier


func main() {
	//命令行解析路径
	filename := pflag.Identify()

	//解析json文件路径，解析数据
	JsonPare := loadpath.NewJsonStruct()
	vmap := Vmap{}
    ////下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
    JsonPare.Load(filename,&vmap)  //经过此方法数据存入到map v中

	//解析ini文件，调出监听端口
	port := RandIni()
	// 创建默认的路由引擎
	r:=gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message" : "hello",
		})
	})

	//输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
	r.GET("/getLicitSoldier", func(c *gin.Context) {
		rarity := c.Query("rarity")
		unK := c.Query("unK")
		cvc := c.Query("cvc")
		soldier := soldiermethod.GetLicitSoldier(vmap, rarity, unK, cvc)
		c.JSON(200,soldier)
	})

	//输入士兵id获取稀有度
	r.GET("/getRarity", func(c *gin.Context) {
		id := c.Query("Id")
		rarity := soldiermethod.GetRarity(vmap, id)
		c.JSON(200,rarity)
	})

	//输入士兵id获取战力
	r.GET("/getCombatPoints", func(c *gin.Context) {
		id := c.Query("Id")
		points := soldiermethod.GetCombatPoints(vmap, id)
		c.JSON(200,points)
	})

	//输入cvc获取所有合法的士兵
	r.GET("/getCvcLicitSoldier", func(c *gin.Context) {
		cvc := c.Query("cvc")
		soldier := soldiermethod.GetCvcLicitSoldier(vmap, cvc)
		c.JSON(200,soldier)
	})

	//获取每个阶段解锁相应士兵的json数据
	r.GET("/getUnkSoldier", func(c *gin.Context) {
		soldier := soldiermethod.GetUnkSoldier(vmap)
		c.JSON(200,soldier)
	})

	r.Run(":"+strconv.Itoa(port))
}

//解析app.ini文件，并调出端口参数
func RandIni() int {
	conf := ini.NewIniFileConfigSource("./app.ini")
	port := conf.GetIntDefault("server.HttpPort",3306)
	return port
}

