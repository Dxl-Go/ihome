package main

import (
	"github.com/gin-gonic/gin"
	"ihomebj5q/model"
	"fmt"
	"ihomebj5q/controller"
)

func main() {
	router := gin.Default()
	/*r1 := router.Group("/v1")
	{
		r1.GET("/efg", func(ctx *gin.Context) {
			ctx.Writer.WriteString("abcdefg")
		})
	}
	r2 := router.Group("/v2")
	{
		r2.GET("/efg", func(ctx *gin.Context) {
			ctx.Writer.WriteString("111222")
		})
	}
	//model.InitModel()
	//model.InsertData()
	//model.SearchData()
	//model.UpdateData()
	//model.DeleteData()
	//model.InsertTeacher()*/
	model.InitRedis()
	err := model.InitDb()
	if err!=nil {
		//把错误打印到日志上,
		fmt.Println(err)
		return
	}

	//路由模块
	//router.Group("/")
	//展示静态页面
	router.Static("/home","view")

	r1:=router.Group("/api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)
		r1.GET("/session",controller.GetSession)
	}



	router.Run(":8099")
}
