package handler

import (
	"context"

	//"github.com/micro/go-micro/util/log"

	getArea "ihomebj5q/service/getArea/proto/getArea"
	"ihomebj5q/service/getArea/model"
	"ihomebj5q/utils"
)

type GetArea struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetArea) MicroGetArea(ctx context.Context, req *getArea.Request, rsp *getArea.Response) error {
	//获取地址信息
	areas, err := model.GetArea()
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	for _, v := range areas {
		var areaInfo getArea.AreaInfo
		areaInfo.Aid = int32(v.Id)
		areaInfo.Aname = v.Name
		rsp.Data=append(rsp.Data,&areaInfo)
	}

	return nil
}
