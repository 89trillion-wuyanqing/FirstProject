package router

import (
	"FirstProject/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {

	engine.POST("/getArmys", ctrl.GetArmys)
	engine.POST("/getArmyRarity", ctrl.GetArmyRarity)
	engine.POST("/getArmyAtk", ctrl.GetArmyAtk)
	engine.POST("/getArmysByCvc", ctrl.GetArmysByCvc)
	engine.GET("/getArmysByStage", ctrl.GetArmysByStage)

}
