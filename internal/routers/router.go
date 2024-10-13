package routers

import "github.com/gin-gonic/gin"

func RootRouter(){
	r := gin.Default()
	r.Run()	
}

func UserRouter(r *gin.Engine){
	r.GET("/users", controllers.GetUser)
}	
	