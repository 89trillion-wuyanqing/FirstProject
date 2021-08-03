package model

type Army struct {
	Id          string `json:"id"`          //士兵id
	UnlockArena string `json:"UnlockArena"` //解锁阶段
	Rarity      string `json:"Rarity"`      //稀有度
	Cvc         string `json:"Cvc"`         //客户端版本
	Atk         string `json:"Atk"`         //战斗力
}
