package ctrl

import (
	"FirstProject/internal/handler"
	"FirstProject/internal/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArmyController struct {
}

func NewArmyController() *ArmyController {
	return &ArmyController{}
}

var armyHandler = new(handler.ArmyHandler)

func (this *ArmyController) GetArmys() gin.HandlerFunc {

	return func(context *gin.Context) {
		rarity, _ := context.GetPostForm("rarity")
		unlockArena, _ := context.GetPostForm("unlockArena")
		cvc, _ := context.GetPostForm("cvc")
		_, e1 := strconv.Atoi(rarity)
		_, e2 := strconv.Atoi(cvc)
		unlockArenaint, e3 := strconv.Atoi(unlockArena)

		if e1 != nil {

			context.JSON(200, model.Result{Code: "205", Msg: "请正确填写rarity参数", Data: nil})
			return
		} else if e2 != nil {

			context.JSON(200, model.Result{Code: "206", Msg: "请正确填写cvc参数", Data: nil})
			return
		} else if e3 != nil {

			context.JSON(200, model.Result{Code: "207", Msg: "请正确填写unlockArena参数", Data: nil})
			return
		}
		newHandler := handler.ArmyHandler{}
		result := newHandler.GetArmys(rarity, unlockArenaint, cvc)
		context.JSON(200, result)

	}

}

func (this *ArmyController) GetArmyRarity() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := context.GetPostForm("id")
		_, e1 := strconv.Atoi(id)

		//rarity, err := handler.GetArmyRarity(id)
		if e1 != nil {
			context.JSON(200, model.Result{Code: "208", Msg: "请正确填写id参数", Data: nil})

		} else {
			result := armyHandler.GetArmyRarity(id)
			context.JSON(200, result)
		}
	}

}

func (this *ArmyController) GetArmyAtk() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := context.GetPostForm("id")
		_, e1 := strconv.Atoi(id)
		//atk, err := handler.GetArmyAtk(id)

		if e1 != nil {
			context.JSON(200, model.Result{Code: "208", Msg: "请正确填写id参数", Data: nil})
		} else {
			result := armyHandler.GetArmyAtk(id)
			context.JSON(200, result)
		}
	}
}

func (this *ArmyController) GetArmysByCvc() gin.HandlerFunc {
	return func(context *gin.Context) {
		cvc, _ := context.GetPostForm("cvc")
		_, e1 := strconv.Atoi(cvc)
		//returnMap, err := handler.GetArmysByCvc(cvc)

		if e1 != nil {
			context.JSON(200, model.Result{Code: "206", Msg: "请正确填写cvc参数", Data: nil})
		} else {
			result := armyHandler.GetArmysByCvc(cvc)
			context.JSON(200, result)
		}
	}
}

func (this *ArmyController) GetArmysByStage() gin.HandlerFunc {

	return func(context *gin.Context) {
		//returnMap, err := handler.GetArmysByStage()
		result := armyHandler.GetArmysByStage()
		context.JSON(200, result)
	}
}
