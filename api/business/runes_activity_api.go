package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/common/api_result"
	"github.com/Pump-Elf-Ranch/per_apollo/common/bigint"
	"github.com/Pump-Elf-Ranch/per_apollo/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func RunesActivityApi(rg *gin.Engine) {
	r := rg.Group("/v1/activity")
	r.GET("list", runesActivityList)
}

func runesActivityList(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	if pageSizeStr == "" || pageNumStr == "" {
		pageSizeStr = "10"
		pageNumStr = "1"
	}
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	activityTypeStr := c.Query("activityType")
	var activityType []int
	if activityTypeStr != "" {
		activityTypeArr := strings.Split(activityTypeStr, ",")
		for _, v := range activityTypeArr {
			activityType = append(activityType, bigint.StringToInt(v))
		}
	}
	infos, count := service.BaseService.RunesActivityService.RunesActivityList(activityType, pageNum, pageSize)
	page := api_result.NewPage(infos, int(count), pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}
