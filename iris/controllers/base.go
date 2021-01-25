package controllers

import (
	"github.com/kataras/iris/v12"
)

type BaseController struct {
}

type OperationResult struct {
	Flag         bool        `json:"flag"`
	ErrCode      string      `json:"errCode"`
	ErrMessage   string      `json:"errMessage"`
	ExceptionMsg string      `json:"exceptionMsg"`
	ResData      interface{} `json:"resData"`
}

func (c *BaseController) EchoJson(ctx iris.Context, msg interface{})  {
	ctx.JSON(OperationResult{
		ExceptionMsg: msg.(string),
	})
	ctx.StopExecution()
}

func (c *BaseController) EchoJsonOk(ctx iris.Context, resData ...interface{}) {
	if resData == nil {
		resData = []interface{}{""}
	}
	ctx.JSON(OperationResult{
		Flag:    true,
		ResData: resData[0],
	})
	ctx.StopExecution()
}

func (c *BaseController) EchoErr(ctx iris.Context, err error)  {
	ctx.JSON(OperationResult{
		Flag:         false,
		ErrCode:      "400",
		ErrMessage:   err.Error(),
	})
	ctx.StopExecution()
}
