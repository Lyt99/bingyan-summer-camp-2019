package Err

var BindingFailedJson = map[string]interface{}{
	"success": false,
	"error":   "错误：缺少必要数据或数据不合法，请检查你的输入",
	"data":    ""}

var UserExistJson = map[string]interface{}{
	"success": false,
	"error":   "错误：用户名已存在",
	"data":    ""}

var InsertFailedJson = map[string]interface{}{
	"success": false,
	"error":   "错误：注册失败，服务器错误，请联系管理员",
	"data":    ""}

var GetFailedJson = map[string]interface{}{
	"success": false,
	"error":   "错误：获取失败",
	"data":    ""}

var NoKeyJson = map[string]interface{}{
	"success": false,
	"error":   "错误：缺少关键词",
	"data":    ""}

var UserNotExistJson = map[string]interface{}{
	"success": false,
	"error":   "错误：用户不存在",
	"data":    ""}

var DeleteFailedJson = map[string]interface{}{
	"success": false,
	"error":   "错误：无删除权限",
	"data":    ""}

var CommodityNotExistJson = map[string]interface{}{
	"success": false,
	"error":   "错误：商品不存在",
	"data":    ""}

var CollectionExistJson = map[string]interface{}{
	"success": false,
	"error":   "错误：收藏已存在",
	"data":    ""}

var IdGetFailedJson = map[string]interface{}{
	"success": false,
	"error":   "错误：用户登录状态读取失败，请尝试重新登录",
	"data":    ""}

var PutPicFailedJson = map[string]interface{}{
	"success": false,
	"error":   "错误：服务器发生错误，请联系管理员",
	"data":    ""}

var PicNotSelectedJson = map[string]interface{}{
	"success": false,
	"error":   "错误：未选择任何图片",
	"data":    ""}

var InvalidMobileJson = map[string]interface{}{
	"success": false,
	"error":   "错误：手机号码格式错误",
	"data":    ""}

var InvalidEmailJson = map[string]interface{}{
	"success": false,
	"error":   "错误：邮箱地址格式错误",
	"data":    ""}