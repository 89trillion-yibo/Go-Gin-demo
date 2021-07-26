package service

import (
	"awesomeProject/gin_demo/internal/model"
	"awesomeProject/gin_demo/utils"
)

var Soldier  = utils.Soldier

type Unkmap map[string][]model.Soldier

//输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func  GetLicitSoldier(rarity string,unK string,cvc string ) map[string]model.Soldier {
	meets := make(map[string]model.Soldier)
	//解锁阶段为空，特殊情况单独判断
	if unK == "" {
		for k,v := range Soldier{
			if v.Cvc == cvc && v.Rarity == rarity && v.UnlockArena == unK {
				meets[k] = v
			}
		}
	}else {  //输出符合其他条件并且小于当前解锁阶段的士兵集合
		for k,v := range Soldier{
			if v.Cvc == cvc && v.Rarity == rarity && v.UnlockArena <= unK {
				meets[k] = v
			}
		}
	}
	return meets
}

//输入士兵id获取稀有度
func GetRarity(id string) string {
	for k,v := range Soldier{
		if k==id{
			return v.Rarity
		}
	}
	return ""
}

//输入士兵id获取战力
func GetCombatPoints(id string) string {
	for k,v := range Soldier{
		if k==id{
			return v.CombatPoints
		}
	}
	return ""
}

//输入cvc获取所有合法的士兵
func GetCvcLicitSoldier(cvc string ) map[string]model.Soldier {
	meets := make(map[string]model.Soldier)
	for k,v := range Soldier{
		if v.Cvc == cvc{
			meets[k] = v
		}
	}
	return meets
}

//获取每个阶段解锁相应士兵的json数据
func GetUnkSoldier() map[string][]model.Soldier {
	unkmap := Unkmap{}
	for k,v := range Soldier{
		//判断map是否含有该key，没有则初始化该key下的士兵切片
		if _,ok := unkmap[v.UnlockArena];!ok{
			unkmap[v.UnlockArena] = make([]model.Soldier,0,len(Soldier))
		}
		//将数据添加到key对应的士兵切片下
		unkmap[v.UnlockArena] = append(unkmap[v.UnlockArena], Soldier[k])
	}
	return unkmap
}
