package soldiermethod

//士兵结构体
type Soldier struct {
	Id           string  //id
	Rarity       string  //稀有度
	UnlockArena  string  //解锁阶段
	CombatPoints string  //战斗力
	Cvc          string  //客户端版本号
}

type Unkmap map[string][]Soldier

//输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func  GetLicitSoldier(vmap map[string]Soldier,rarity string,unK string,cvc string ) map[string]Soldier {
	meets := make(map[string]Soldier)
	//解锁阶段为空，特殊情况单独判断
	if unK == "" {
		for k,v := range vmap{
			if v.Cvc == cvc && v.Rarity == rarity && v.UnlockArena == unK {
				meets[k] = v
			}
		}
	}else {  //输出符合其他条件并且小于当前解锁阶段的士兵集合
		for k,v := range vmap{
			if v.Cvc == cvc && v.Rarity == rarity && v.UnlockArena <= unK {
				meets[k] = v
			}
		}
	}
	return meets
}

//输入士兵id获取稀有度
func GetRarity(vmap map[string]Soldier,id string) string {
	for k,v := range vmap{
		if k==id{
			return v.Rarity
		}
	}
	return "无此士兵id"
}

//输入士兵id获取战力
func GetCombatPoints(vmap map[string]Soldier,id string) string {
	for k,v := range vmap{
		if k==id{
			return v.CombatPoints
		}
	}
	return "无此士兵id"
}

//输入cvc获取所有合法的士兵
func GetCvcLicitSoldier(vmap map[string]Soldier,cvc string ) map[string]Soldier {
	meets := make(map[string]Soldier)
	for k,v := range vmap{
		if v.Cvc == cvc{
			meets[k] = v
		}
	}
	return meets
}

//获取每个阶段解锁相应士兵的json数据
func GetUnkSoldier(vmap map[string]Soldier) map[string][]Soldier {
	unkmap := Unkmap{}
	for k,v := range vmap{
		//判断map是否含有该key，没有则初始化该key下的士兵切片
		if _,ok := unkmap[v.UnlockArena];!ok{
			unkmap[v.UnlockArena] = make([]Soldier,0,len(vmap))
		}
		//将数据添加到key对应的士兵切片下
		unkmap[v.UnlockArena] = append(unkmap[v.UnlockArena], vmap[k])
	}
	return unkmap
}
