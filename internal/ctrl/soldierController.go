package ctrl

import (
	"awesomeProject/gin_demo/internal/service"
	"awesomeProject/gin_demo/internal/soldierEorrer"
	"github.com/gin-gonic/gin"
	"net/http"
)

//输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func LicitSol(c *gin.Context) {
	rarity,a := c.GetQuery("rarity")
	unK,b := c.GetQuery("unK")
	cvc,e := c.GetQuery("cvc")
	if !a || !b || !e || rarity == "" || cvc == ""{
		//参数为空校验
		c.JSON(http.StatusBadRequest,soldierEorrer.Parameters)
	}else {
		//参数不为空返回数据
		soldier := service.GetLicitSoldier(rarity, unK, cvc)
		if len(soldier) == 0{
			c.JSON(http.StatusOK,soldierEorrer.NoData)
		}else {
			c.JSON(http.StatusOK,soldierEorrer.OK.AddData(soldier))
		}
	}
}

////输入士兵id获取稀有度
func RaritySol(c *gin.Context)  {
	id, a := c.GetQuery("Id")
	if !a || id == "" {
		c.JSON(http.StatusBadRequest,soldierEorrer.Parameters)
	}else {
		rarity := service.GetRarity(id)
		if rarity == ""{
			c.JSON(http.StatusOK,soldierEorrer.NoData)
		}else {
			c.JSON(http.StatusOK,soldierEorrer.OK.AddData(rarity))
		}
	}
}

//输入士兵id获取战力
func CombatPointsSol(c *gin.Context)  {
	id, a := c.GetQuery("Id")
	if !a || id == "" {
		c.JSON(http.StatusBadRequest,soldierEorrer.Parameters)
	}else {
		points := service.GetCombatPoints(id)
		if points == ""{
			c.JSON(http.StatusOK,soldierEorrer.NoData)
		}else {
			c.JSON(http.StatusOK,soldierEorrer.OK.AddData(points))
		}
	}
}


//输入cvc获取合法的士兵
func CvcLicitSol(c *gin.Context)  {
	cvc, a := c.GetQuery("cvc")
	if !a || cvc == "" {
		c.JSON(http.StatusBadRequest,soldierEorrer.Parameters)
	}else {
		soldier := service.GetCvcLicitSoldier(cvc)
		if len(soldier) == 0{
			c.JSON(http.StatusOK,soldierEorrer.NoData)
		}else {
			c.JSON(http.StatusOK,soldierEorrer.OK.AddData(soldier))
		}
	}
}

//获取每个阶段解锁相应士兵的json数据
func UnkSol(c *gin.Context) {
	soldier := service.GetUnkSoldier()
	if len(soldier) == 0{
		c.JSON(http.StatusOK,soldierEorrer.NoData)
	}else {
		c.JSON(http.StatusOK,soldierEorrer.OK.AddData(soldier))
	}
}