package controllers

import (
	"github.com/HFO4/cloudreve/pkg/serializer"
	"github.com/HFO4/cloudreve/pkg/util"
	"github.com/HFO4/cloudreve/service/callback"
	"github.com/gin-gonic/gin"
)

// RemoteCallback 远程上传回调
func RemoteCallback(c *gin.Context) {
	var callbackBody callback.RemoteUploadCallbackService
	if err := c.ShouldBindJSON(&callbackBody); err == nil {
		res := callback.ProcessCallback(callbackBody, c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// QiniuCallback 七牛上传回调
func QiniuCallback(c *gin.Context) {
	var callbackBody callback.UploadCallbackService
	if err := c.ShouldBindJSON(&callbackBody); err == nil {
		res := callback.ProcessCallback(callbackBody, c)
		if res.Code != 0 {
			c.JSON(401, serializer.QiniuCallbackFailed{Error: res.Msg})
		} else {
			c.JSON(200, res)
		}
	} else {
		c.JSON(401, ErrorResponse(err))
	}
}

// OSSCallback 阿里云OSS上传回调
func OSSCallback(c *gin.Context) {
	var callbackBody callback.UploadCallbackService
	if err := c.ShouldBindJSON(&callbackBody); err == nil {
		if callbackBody.PicInfo == "," {
			callbackBody.PicInfo = ""
		}
		res := callback.ProcessCallback(callbackBody, c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpyunCallback 又拍云上传回调
func UpyunCallback(c *gin.Context) {
	var callbackBody callback.UpyunCallbackService
	if err := c.ShouldBind(&callbackBody); err == nil {
		if callbackBody.Code != 200 {
			util.Log().Debug(
				"又拍云回调返回错误代码%d，信息：%s",
				callbackBody.Code,
				callbackBody.Message,
			)
			return
		}
		res := callback.ProcessCallback(callbackBody, c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}