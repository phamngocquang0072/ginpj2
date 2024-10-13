package main

import (
	// "github.com/gin-gonic/gin"
	"github.com/phamngocquang0072/ginpj2/initializers"
	"github.com/phamngocquang0072/ginpj2/internal/routers"
)

func init() {
	// loadEnv()
	initializers.loadEnv()
	initializers.ConnectDB()
	routers.RootRouter()
}

func main() {
	
}