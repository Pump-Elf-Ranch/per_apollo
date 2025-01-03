package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/common/api_result"
	"github.com/Pump-Elf-Ranch/per_apollo/common/bigint"
	"github.com/Pump-Elf-Ranch/per_apollo/common/enum"
	"github.com/Pump-Elf-Ranch/per_apollo/service"
	"github.com/gin-gonic/gin"
)

func RunesListedApi(rg *gin.Engine) {
	r := rg.Group("/v1/listed")
	r.GET("list", runesListedList)
}

func runesListedList(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")
	if pageSizeStr == "" || pageNumStr == "" {
		pageSizeStr = "10"
		pageNumStr = "1"
	}
	pageSize := bigint.StringToInt(pageSizeStr)
	pageNum := bigint.StringToInt(pageNumStr)
	orderByType := c.Query("orderByType")
	if orderByType != "" {
		byTypeBool := checkOrderByType(orderByType)
		if !byTypeBool {
			api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
			return
		}
	}
	tokenAddress := c.Query("tokenAddress")
	orderBy := c.Query("orderBy")
	if orderBy != "" {
		orderByBool := checkOrderBy(orderByType)
		if !orderByBool {
			api_result.NewApiResult(c).Error(enum.ParamErr.Code, enum.ParamErr.Msg)
			return
		}
	} else {
		orderBy = "desc"
	}
	infos, count := service.BaseService.RunesListedService.RunesListedList(orderByType, orderBy, tokenAddress, pageNum, pageSize)
	page := api_result.NewPage(infos, int(count), pageNum, pageSize)
	api_result.NewApiResult(c).Success(page)
}

func checkOrderByType(orderByType string) bool {
	if orderByType == "total" || orderByType == "price" || orderByType == "create_time" {
		return true
	}
	return false
}

func checkOrderBy(orderBy string) bool {
	if orderBy == "asc" || orderBy == "desc" {
		return true
	}
	return false
}
