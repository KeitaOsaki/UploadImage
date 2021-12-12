package pkg

import (
"file-server/pkg/controller"
"github.com/gin-gonic/gin"
"github.com/gin-contrib/cors"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {

	Server = gin.Default()
	//allows all origins
	Server.Use(cors.Default())

	Server.POST("/upload", controller.Upload())

}
