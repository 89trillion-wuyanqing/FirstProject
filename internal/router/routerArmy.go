package router

import (
	"FirstProject/internal/ctrl"
	"github.com/gin-gonic/gin"
)

type RouterArmys struct {
}

func (r *RouterArmys) Router(engine *gin.Engine) {
	ctrlArmy := ctrl.NewArmyController()
	engine.POST("/getArmys", ctrlArmy.GetArmys())
	engine.POST("/getArmyRarity", ctrlArmy.GetArmyRarity())
	engine.POST("/getArmyAtk", ctrlArmy.GetArmyAtk())
	engine.POST("/getArmysByCvc", ctrlArmy.GetArmysByCvc())
	engine.GET("/getArmysByStage", ctrlArmy.GetArmysByStage())

}
