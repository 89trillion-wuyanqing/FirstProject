package ctrl

import (
	"FirstProject/internal/handler"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*type ArmyController struct {
	*gin.Engine
}

// 这里是构造函数
func NewArmyController(e *gin.Engine) *ArmyController {
	return &ArmyController{e}
}*/

func GetArmys(context *gin.Context) {

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

	returnMap, err := handler.GetArmys(rarity, unlockArenaint, cvc)
	if errmsg != "" {
		context.JSON(200, map[string]interface{}{
			"code":    201,
			"message": "ERROR",
			"data":    errmsg,
		})
	} else if err != nil {
		context.JSON(200, map[string]interface{}{
			"code":    201,
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

func GetArmyRarity(context *gin.Context) {
	id, _ := context.GetPostForm("id")
	_, e1 := strconv.Atoi(id)

	rarity, err := handler.GetArmyRarity(id)
	if e1 != nil {
		context.JSON(200, map[string]interface{}{
			"code":    201,
			"message": "ERROR",
			"data":    "请正确填写id参数 ",
		})
	} else if err != nil {
		context.JSON(200, map[string]interface{}{
			"code":    201,
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

func GetArmyAtk(context *gin.Context) {
	id, _ := context.GetPostForm("id")
	_, e1 := strconv.Atoi(id)
	atk, err := handler.GetArmyAtk(id)
	if e1 != nil {
		context.JSON(200, map[string]interface{}{
			"code":    201,
			"message": "ERROR",
			"data":    "请正确填写id参数 ",
		})
	} else if err != nil {
		context.JSON(200, map[string]interface{}{
			"code":    201,
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

func GetArmysByCvc(context *gin.Context) {
	cvc, _ := context.GetPostForm("cvc")
	_, e1 := strconv.Atoi(cvc)
	returnMap, err := handler.GetArmysByCvc(cvc)
	if e1 != nil {
		context.JSON(200, map[string]interface{}{
			"code":    201,
			"message": "ERROR",
			"data":    "请正确填写cvc参数 ",
		})
	} else if err != nil {
		context.JSON(200, map[string]interface{}{
			"code":    201,
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

func GetArmysByStage(context *gin.Context) {

	returnMap, err := handler.GetArmysByStage()
	if err != nil {
		context.JSON(200, map[string]interface{}{
			"code":    201,
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
