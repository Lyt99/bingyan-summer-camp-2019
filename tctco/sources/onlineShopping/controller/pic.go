package controller

import (
	"github.com/gin-gonic/gin"
	"onlineShopping/model"
	"onlineShopping/settings"
	"strings"
)

func UploadPic(c *gin.Context) {
	response := ResponseInit()
	file, err := c.FormFile("file")
	if err != nil {
		response["error"] = "文件上传不成功"
		c.JSON(400, response)
		return
	}

	if !fmtCheckPic(file.Filename) {
		response["error"] = "图片格式不正确"
		c.JSON(400, response)
		return
	}

	err = model.OSSavePic(file)
	if err != nil {
		internalError(c)
		return
	}

	c.JSON(200, gin.H{"url": getImageUrl(file.Filename)})
}

func fmtCheckPic(filename string) bool {
	if strings.HasSuffix(filename, ".jpg") || strings.HasSuffix(filename, ".png") || strings.HasSuffix(filename, ".bmp") || strings.HasSuffix(filename, ".jpeg") {
		return true
	}
	return false
}

//this func generates a url, through which users
//can access the image
func getImageUrl(filename string) string {
	return settings.ImagePrefixUrl + settings.ImageSavePath + filename
}
