package Err

var BindingFailedJson = map[string]interface{}{
	"success": false,
	"error":   "缺少必要数据或数据不合法",
	"data":    ""}

var UserExistJson = map[string]interface{}{
	"success": false,
	"error":   "用户名已存在",
	"data":    ""}

var InsertFailedJson = map[string]interface{}{
	"success": false,
	"error":   "注册失败：服务器错误",
	"data":    ""}

var GetFailedJson = map[string]interface{}{
	"success": false,
	"error":   "获取失败",
	"data":    ""}

var NoKeyJson = map[string]interface{}{
	"success": false,
	"error":   "缺少关键词",
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

var IdGetFailedJson = map[string]interface{}{
	"success": false,
	"error":   "用户登录状态读取失败：请尝试重新登录",
	"data":    ""}
