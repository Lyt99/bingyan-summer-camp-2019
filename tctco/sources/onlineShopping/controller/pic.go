package controller

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

	src, err := file.Open()
	if err != nil {
		internalError(c)
		return
	}
	file.Filename = hashFileName(file.Filename)
	dst, err := os.Create(filepath.Join("./upload", filepath.Base(file.Filename)))
	if err != nil {
		internalError(c)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
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

func hashFileName(filename string) string {
	m := md5.New()
	filename = filename
	m.Write([]byte(filename))
	extension := filepath.Ext(filename)

	return hex.EncodeToString(m.Sum(nil)) + strconv.FormatInt(time.Now().Unix(), 10) + extension
}

func getImageUrl(filename string) string {
	return ImagePrefixUrl + ImageSavePath + filename
}
