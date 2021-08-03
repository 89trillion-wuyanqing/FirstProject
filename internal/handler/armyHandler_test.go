package handler

import "testing"

func TestGetArmyAtk(t *testing.T) {
	armyhandler := ArmyHandler{}
	atk, err := armyhandler.GetArmyAtk("10502")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("id为10502的士兵战斗力为" + atk)
}

func TestGetArmys(t *testing.T) {
	armyhandler := ArmyHandler{}
	returnmap, err := armyhandler.GetArmys("2", 2, "1000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(returnmap)
}

func TestGetArmyRarity(t *testing.T) {
	armyhandler := ArmyHandler{}
	rarity, err := armyhandler.GetArmyRarity("10501")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("id为10501的士兵的稀有度为" + rarity)
}

func TestGetArmysByCvc(t *testing.T) {
	armyhandler := ArmyHandler{}
	returnmap, err := armyhandler.GetArmysByCvc("1000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(returnmap)
}
func TestGetArmysByStage(t *testing.T) {
	armyhandler := ArmyHandler{}
	returnmap, err := armyhandler.GetArmysByStage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(returnmap)
}
