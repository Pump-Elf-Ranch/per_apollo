package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/common/api_result"
	"github.com/Pump-Elf-Ranch/per_apollo/common/bigint"
	"github.com/Pump-Elf-Ranch/per_apollo/common/enum"
	"github.com/Pump-Elf-Ranch/per_apollo/service"
	"github.com/gin-gonic/gin"
)

func MintListedApi(rg *gin.Engine) {
	r := rg.Group("/v1/mint")
	r.GET("list", mint_list)
}

func mint_list(c *gin.Context) {
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

	infos, count := service.BaseService.MintListedService.MintListedList(userAddress, pageNum, pageSize)
	page := api_result.NewPage(infos, int(count), pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}
