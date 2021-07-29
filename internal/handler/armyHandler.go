package handler

import (
	"FirstProject/internal/model"
	"FirstProject/internal/utils"
	"errors"
	"strconv"
)

//

func GetArmys(rarity string, unlockArena int, cvc string) (map[string]model.Army, error) {
	var returnMap map[string]model.Army
	returnMap = make(map[string]model.Army)

	army, e := utils.GetNewJson()

	if e != nil {
		return nil, e
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
						return nil, e

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

	return returnMap, nil

}

func GetArmyRarity(id string) (string, error) {

	army, e := utils.GetNewJson()
	if e != nil {
		return "", e
	}

	for _, v := range army {
		//cvc匹配
		if v.Id == id {
			return v.Rarity, nil

		}

	}
	return "", errors.New("id没有匹配到士兵")
}

func GetArmyAtk(id string) (string, error) {

	army, e := utils.GetNewJson()
	if e != nil {
		return "", e
	}

	for _, v := range army {
		//cvc匹配
		if v.Id == id {
			return v.Atk, nil

		}

	}
	return "", errors.New("id没有匹配到士兵")
}

func GetArmysByCvc(cvc string) (map[string]model.Army, error) {
	var returnMap map[string]model.Army
	returnMap = make(map[string]model.Army)
	army, e := utils.GetNewJson()
	if e != nil {
		return nil, e
	}

	for k, v := range army {
		//cvc匹配
		if v.Cvc == cvc {
			returnMap[k] = v

		}

	}
	return returnMap, nil
}

func GetArmysByStage() (map[string][]model.Army, error) {
	var returnMap map[string][]model.Army
	returnMap = make(map[string][]model.Army)
	army, e := utils.GetNewJson()
	if e != nil {
		return nil, e
	}

	for _, v := range army {
		if v.UnlockArena == "" {
			returnMap["0"] = append(returnMap["0"], v)
		} else {
			returnMap[v.UnlockArena] = append(returnMap[v.UnlockArena], v)
		}
	}
	return returnMap, nil

}
