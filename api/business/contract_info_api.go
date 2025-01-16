package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/common/api_result"
	"github.com/Pump-Elf-Ranch/per_apollo/service"
	"github.com/gin-gonic/gin"
)

func ContractInfoApi(rg *gin.Engine) {
	r := rg.Group("/v1/contract")
	r.GET("info", info)
}

func info(c *gin.Context) {
	infos := service.BaseService.ContractInfoService.ContractInfos()
	api_result.NewApiResult(c).Success(infos)
}
