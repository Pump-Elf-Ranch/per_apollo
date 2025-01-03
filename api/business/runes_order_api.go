package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/common/api_result"
	"github.com/Pump-Elf-Ranch/per_apollo/common/bigint"
	"github.com/Pump-Elf-Ranch/per_apollo/common/enum"
	"github.com/Pump-Elf-Ranch/per_apollo/service"
	"github.com/gin-gonic/gin"
)

func RunesOrderApi(rg *gin.Engine) {
	r := rg.Group("/v1/order")
	r.GET("list", orderList)
	r.GET("info", info)
}

func orderList(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	if pageSizeStr == "" || pageNumStr == "" {
		pageSizeStr = "10"
		pageNumStr = "1"
	}
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	userAddress := c.Query("userAddress")
	if userAddress == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}
	tokenAddress := c.Query("tokenAddress")
	orderId := c.Query("orderId")
	infos, count := service.BaseService.RunesOrderService.RunesOrderList(userAddress, tokenAddress, orderId, pageNum, pageSize)
	page := api_result.NewPage(infos, int(count), pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

func info(c *gin.Context) {
	orderId := c.Query("orderId")

	if orderId == "" {
		api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
		return
	}
	resultInfo := service.BaseService.RunesOrderService.RunesOrderInfo(orderId)
	api_result.NewApiResult(c).Success(resultInfo)
}
