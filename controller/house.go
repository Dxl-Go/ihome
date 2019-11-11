package controller

import (
	"github.com/gin-gonic/gin"
	getArea "ihomebj5q/proto/getArea/proto/getArea"
	//"github.com/micro/go-micro/client"
	"context"
	"fmt"
	"net/http"
	//"github.com/micro/go-micro/registry/consul"
	//"github.com/micro/go-micro"
	//"ihomebj5q/utils"
	//"github.com/micro/go-micro/client"
	"ihomebj5q/utils"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro"
)

func GetArea(ctx *gin.Context) {
	/*resp:=make(map[string]interface{})
	defer ctx.JSON(http.StatusOK,resp)
	//获取数据库数据
	areas,err:=model.GetArea()
	if err!=nil {
		fmt.Println("获取所有地域信息错误")
		resp["errno"]=utils.RECODE_DBERR
		resp["errmsg"]=utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	resp["errno"]=utils.RECODE_OK
	resp["errmsg"]=utils.RecodeText(utils.RECODE_OK)
	resp["data"]=areas*/


	//从consul中获取服务
	consulRegistry:=consul.NewRegistry()
	micService:=micro.NewService(
		micro.Registry(consulRegistry),
	)
	microClient:=getArea.NewGetAreaService("go.micro.srv.getArea",micService.Client())
	//调用远程服务
	resp,err:=microClient.MicroGetArea(context.TODO(),&getArea.Request{})
	if err!=nil {
		fmt.Println(err)
	}
ctx.JSON(http.StatusOK,resp)
}
func GetSession(ctx *gin.Context)  {
	resp:=make(map[string]interface{})
	resp["errno"]=utils.RECODE_LOGINERR
	resp["errmsg"]=utils.RecodeText(utils.RECODE_LOGINERR)
	ctx.JSON(http.StatusOK,resp)
}
