package Err

import "errors"
var BindingFailed=errors.New("缺少必要数据或数据不合法")
var BindingFailedJson=map[string]interface{}{
	"success": false,
	"error":"缺少必要数据或数据不合法",
	"data":""}

var UserExistJson =map[string]interface{}{
	"success": false,
	"error":"用户已存在",
	"data":""}

var InsertFailedJson=map[string]interface{}{
	"success": false,
	"error":"服务器错误",
	"data":""}

var UserNotExist=errors.New("用户不存在")
var UserNotExistJson=map[string]interface{}{
	"success": false,
	"error":"用户不存在",
	"data":""}

var WrongPsw=errors.New("密码错误")
var WrongPswJson=map[string]interface{}{
	"success": false,
	"error":"密码错误",
	"data":""}

var GetFailedJson=map[string]interface{}{
	"success": false,
	"error":"获取失败",
	"data":""}