package controller

import (
	"V2RayA/model/v2ray"
	"V2RayA/persistence/configure"
	"V2RayA/service"
	"V2RayA/tools"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetSiteDatFiles(ctx *gin.Context) {
	tools.ResponseSuccess(ctx, gin.H{"siteDatFiles": service.GetSiteDatFiles()})
}
func GetCustomPac(ctx *gin.Context) {
	tools.ResponseSuccess(ctx, gin.H{
		"customPac":          configure.GetCustomPacNotNil(),
		"V2RayLocationAsset": v2ray.GetV2rayLocationAsset(),
	})
}
func PutCustomPac(ctx *gin.Context) {
	var data struct {
		CustomPac configure.CustomPac `json:"customPac"`
	}
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		tools.ResponseError(ctx, errors.New("参数有误"+err.Error()))
		return
	}
	err = configure.SetCustomPac(&data.CustomPac)
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	tools.ResponseSuccess(ctx, nil)
}