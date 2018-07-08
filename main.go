package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gogo/routes"
)

func main(){
	//numCPU := runtime.NumCPU()
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger())
	routes.RouteRegister(r)
	run(r)
}

func run(r *gin.Engine)  {
	r.Run(":8000")
}

func debug(a interface{})  {
	fmt.Println(a)
}