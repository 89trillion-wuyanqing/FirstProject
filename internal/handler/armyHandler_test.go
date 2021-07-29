package handler

import "testing"

func TestGetArmyAtk(t *testing.T) {

	atk, err := GetArmyAtk("10502")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("id为10502的士兵战斗力为" + atk)
}

func TestGetArmys(t *testing.T) {
	returnmap, err := GetArmys("2", 2, "1000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(returnmap)
}

func TestGetArmyRarity(t *testing.T) {
	rarity, err := GetArmyRarity("10501")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("id为10501的士兵的稀有度为" + rarity)
}

func TestGetArmysByCvc(t *testing.T) {
	returnmap, err := GetArmysByCvc("1000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(returnmap)
}
func TestGetArmysByStage(t *testing.T) {
	returnmap, err := GetArmysByStage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(returnmap)
}
