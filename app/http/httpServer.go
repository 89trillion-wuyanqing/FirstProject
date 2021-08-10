package http

import (
	"FirstProject/internal/router"
	"FirstProject/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// 初始化函数
func Init() {
	// 作为Server的构造器
	engine := gin.Default()

	var r = &router.RouterArmys{}
	r.Router(engine)

	httpPort := utils.GetVal("server", "HttpPort")
	fmt.Println("端口号：" + httpPort)
	var sss = ":" + httpPort
	fmt.Println("端口号" + sss)

	err := engine.Run(sss)
	if err != nil {
		log.Fatal("http服务启动错误" + err.Error())
	}

}
