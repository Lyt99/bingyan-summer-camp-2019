package Err

import "errors"

var WrongPsw = errors.New("密码错误")
var BindingFailed = errors.New("缺少必要数据或数据不合法")
var UserNotExist = errors.New("用户不存在")

var BindingFailedJson = map[string]interface{}{
	"success": false,
	"error":   "缺少必要数据或数据不合法",
	"data":    ""}

var UserExistJson = map[string]interface{}{
	"success": false,
	"error":   "用户已存在",
	"data":    ""}

var InsertFailedJson = map[string]interface{}{
	"success": false,
	"error":   "服务器错误",
	"data":    ""}

var GetFailedJson = map[string]interface{}{
	"success": false,
	"error":   "获取失败",
	"data":    ""}

var NoKeyJson = map[string]interface{}{
	"success": false,
	"error":   "缺少搜索关键词",
	"data":    ""}

var UserNotExistJson = map[string]interface{}{
	"success": false,
	"error":   "用户不存在",
	"data":    ""}

var DeleteFailedJson = map[string]interface{}{
	"success": false,
	"error":   "无删除权限",
	"data":    ""}

var CommodityNotExistJson = map[string]interface{}{
	"success": false,
	"error":   "商品不存在",
	"data":    ""}

var CollectionExistJson = map[string]interface{}{
	"success": false,
	"error":   "收藏已存在",
	"data":    ""}
