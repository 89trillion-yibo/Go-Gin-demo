package router

import (
	"awesomeProject/gin_demo/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	//输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
	r.GET("/getLicitSoldier",ctrl.LicitSol)
	//输入士兵id获取稀有度
	r.GET("/getRarity",ctrl.RaritySol)
	//输入士兵id获取战力
	r.GET("/getCombatPoints",ctrl.CombatPointsSol)
	//输入cvc获取合法的士兵
	r.GET("/getCvcLicitSoldier",ctrl.CvcLicitSol)
	//获取每个阶段解锁相应士兵的json数据
	r.GET("/getUnkSoldier",ctrl.UnkSol)
}
