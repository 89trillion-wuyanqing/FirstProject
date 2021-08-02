package ctrl

import (
	"FirstProject/internal/handler"
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
		var errmsg string = ""
		if e1 != nil {
			errmsg += "请正确填写rarity参数 "
		} else if e2 != nil {
			errmsg += "请正确填写cvc参数 "
		} else if e3 != nil {
			errmsg += "请正确填写unlockArena参数 "
		}
		newHandler := handler.ArmyHandler{}

		returnMap, err := newHandler.GetArmys(rarity, unlockArenaint, cvc)
		if errmsg != "" {
			context.JSON(200, map[string]interface{}{
				"code":    201,
				"message": "ERROR",
				"data":    errmsg,
			})
		} else if err != nil {
			context.JSON(200, map[string]interface{}{
				"code":    202,
				"message": "ERROR",
				"data":    err,
			})
		} else {
			context.JSON(200, map[string]interface{}{
				"code":    200,
				"message": "OK",
				"data":    returnMap,
			})
		}
	}

}

func (this *ArmyController) GetArmyRarity() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := context.GetPostForm("id")
		_, e1 := strconv.Atoi(id)
		rarity, err := armyHandler.GetArmyRarity(id)
		//rarity, err := handler.GetArmyRarity(id)
		if e1 != nil {
			context.JSON(200, map[string]interface{}{
				"code":    201,
				"message": "ERROR",
				"data":    "请正确填写id参数 ",
			})
		} else if err != nil {
			context.JSON(200, map[string]interface{}{
				"code":    202,
				"message": "ERROR",
				"data":    err,
			})
		} else {
			context.JSON(200, map[string]interface{}{
				"code":    200,
				"message": "OK",
				"data":    rarity,
			})
		}
	}

}

func (this *ArmyController) GetArmyAtk() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, _ := context.GetPostForm("id")
		_, e1 := strconv.Atoi(id)
		//atk, err := handler.GetArmyAtk(id)
		atk, err := armyHandler.GetArmyAtk(id)
		if e1 != nil {
			context.JSON(200, map[string]interface{}{
				"code":    201,
				"message": "ERROR",
				"data":    "请正确填写id参数 ",
			})
		} else if err != nil {
			context.JSON(200, map[string]interface{}{
				"code":    202,
				"message": "ERROR",
				"data":    err,
			})
		} else {
			context.JSON(200, map[string]interface{}{
				"code":    200,
				"message": "OK",
				"data":    atk,
			})
		}
	}
}

func (this *ArmyController) GetArmysByCvc() gin.HandlerFunc {
	return func(context *gin.Context) {
		cvc, _ := context.GetPostForm("cvc")
		_, e1 := strconv.Atoi(cvc)
		//returnMap, err := handler.GetArmysByCvc(cvc)
		returnMap, err := armyHandler.GetArmysByCvc(cvc)
		if e1 != nil {
			context.JSON(200, map[string]interface{}{
				"code":    201,
				"message": "ERROR",
				"data":    "请正确填写cvc参数 ",
			})
		} else if err != nil {
			context.JSON(200, map[string]interface{}{
				"code":    202,
				"message": "ERROR",
				"data":    err,
			})
		} else {
			context.JSON(200, map[string]interface{}{
				"code":    200,
				"message": "OK",
				"data":    returnMap,
			})
		}
	}
}

func (this *ArmyController) GetArmysByStage() gin.HandlerFunc {

	return func(context *gin.Context) {
		//returnMap, err := handler.GetArmysByStage()
		returnMap, err := armyHandler.GetArmysByStage()
		if err != nil {
			context.JSON(200, map[string]interface{}{
				"code":    202,
				"message": "ERROR",
				"data":    err,
			})
		} else {
			context.JSON(200, map[string]interface{}{
				"code":    200,
				"message": "OK",
				"data":    returnMap,
			})
		}
	}
}
