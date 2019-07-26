package main

import (
	"fmt"
	"log"
	"path"
	"strings"
)

func main() {
	fullFilename := "https://demo-1258020847.cos.ap-chengdu.myqcloud.com/bkrol24fndiu3unjkqjg.jpg"
	fmt.Println("fullFilename =", fullFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀
	log.Println(len(filenameWithSuffix))
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	fmt.Println("fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
	fmt.Println("filenameOnly =", filenameOnly)
}
