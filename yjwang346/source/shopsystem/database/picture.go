package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ImagePrefixUrl = "http://127.0.0.1:8080"
const ImageSavePath = "/upload/"

/*并不是完全懂*/

//request
//表单，file为图片文件，name为图片文件名
//上传图片，并返回图片地址“
func Picture(c *gin.Context) {//就不返回string了
	name := c.PostForm("name")	//根据name属性获取文件名
	fmt.Println(name)
	//获取文件
	file,err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest,"a Bad request")
	}
	//获取文件名
	filename := file .Filename
	fmt.Println("==============",filename)
	//保存到服务器
	if err := c.SaveUploadedFile(file,file.Filename); err != nil {
		c.String(http.StatusBadRequest,fmt.Sprintf("upload file err: %s",err.Error()))
		return
	}
	c.String(http.StatusOK,"upload successfully")
	//生成url
	generateurl := ImageUrl(filename)
	c.JSON(http.StatusBadRequest, gin.H{
		"success":true,
		"error": "",
		"data":generateurl,//失败的时候留空
	})
	return
}

func ImageUrl(filename string) string {
	return ImagePrefixUrl + ImageSavePath + filename
}