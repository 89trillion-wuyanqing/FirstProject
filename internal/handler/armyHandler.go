package handler

import (
	"FirstProject/internal/model"
	"FirstProject/internal/utils"

	"strconv"
)

//

type ArmyHandler struct {
}

/**
根据稀有度，解锁程度，客户端版本获取符合要求的士兵信息
*/
func (this *ArmyHandler) GetArmys(rarity string, unlockArena int, cvc string) model.Result {
	var returnMap map[string]model.Army
	returnMap = make(map[string]model.Army)

	army, e := utils.GetNewJson()

	if e != nil {
		return model.Result{Code: "201", Msg: "获取不到json文件中士兵信息", Data: nil}
	}

	for k, v := range army {
		//cvc匹配
		if v.Cvc == cvc {
			//稀有度匹配
			if v.Rarity == rarity {
				//解锁程度 如果小于该字段都是以解锁
				if v.UnlockArena != "" {
					ulockInt, e := strconv.Atoi(v.UnlockArena)
					if e != nil {
						return model.Result{Code: "202", Msg: "json文件中解锁阶段的数据格式错误", Data: nil}

					}
					if ulockInt <= unlockArena {
						returnMap[k] = v
					}
				} else {
					returnMap[k] = v
				}
			}

		}

	}

	return model.Result{Code: "200", Msg: "成功", Data: returnMap}

}

/**
根据id获取士兵的稀有度
*/
func (this *ArmyHandler) GetArmyRarity(id string) model.Result {
	var nilArmy model.Army
	army, e := utils.GetNewJson()
	if e != nil {
		return model.Result{Code: "201", Msg: "获取不到json文件中士兵信息", Data: nil}
	}

	if army[id] != nilArmy {
		return model.Result{Code: "200", Msg: "成功", Data: army[id].Rarity}
	}

	return model.Result{Code: "203", Msg: "不存在该士兵", Data: nil}
}

/**
根据士兵id获取士兵攻击力
*/
func (this *ArmyHandler) GetArmyAtk(id string) model.Result {
	var nilArmy model.Army
	army, e := utils.GetNewJson()
	if e != nil {
		return model.Result{Code: "201", Msg: "获取不到json文件中士兵信息", Data: nil}
	}

	if army[id] != nilArmy {
		return model.Result{Code: "200", Msg: "成功", Data: army[id].Atk}
	}

	return model.Result{Code: "203", Msg: "不存在该士兵", Data: nil}
}

/**
根据客户端版本获取符合要求的士兵信息
*/
func (this *ArmyHandler) GetArmysByCvc(cvc string) model.Result {
	var returnMap map[string]model.Army
	returnMap = make(map[string]model.Army)
	army, e := utils.GetNewJson()
	if e != nil {
		return model.Result{Code: "201", Msg: "获取不到json文件中士兵信息", Data: nil}
	}

	for k, v := range army {
		//cvc匹配
		if v.Cvc == cvc {
			returnMap[k] = v

		}

	}
	if len(returnMap) > 0 {
		return model.Result{Code: "200", Msg: "成功", Data: returnMap}
	}

	return model.Result{Code: "204", Msg: "对应客户端版本的士兵信息不存在", Data: nil}
}

/**
将json的格式改变，将其转换成 解锁阶段--士兵
*/
func (this *ArmyHandler) GetArmysByStage() model.Result {
	var returnMap map[string][]model.Army
	returnMap = make(map[string][]model.Army)
	army, e := utils.GetNewJson()
	if e != nil {
		return model.Result{Code: "201", Msg: "获取不到json文件中士兵信息", Data: nil}
	}

	for _, v := range army {
		if v.UnlockArena == "" {
			returnMap["0"] = append(returnMap["0"], v)
		} else {
			returnMap[v.UnlockArena] = append(returnMap[v.UnlockArena], v)
		}
	}
	return model.Result{Code: "200", Msg: "成功", Data: returnMap}

}
