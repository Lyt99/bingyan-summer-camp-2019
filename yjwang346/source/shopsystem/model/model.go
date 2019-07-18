package model

//用户登录以及注册时候用的
type Register struct{
	Id int		`json:"id" form:"id" binding:"required"`
	UserName string		`json:"username" form:"username" binding:"required"`
	Password string		`json:"password" form:"password" binding:"required"`
	Telephone string		`json:"telephone" form:"telephone" binding:"required"`
	Email string	`json:"email" form:"email" binding:"required"`
}

//用户信息页的内容，Visittime在什么时候要进行初始化
type Userinfo struct{
	Id int		`json:"id" form:"id" binding:"required"`
	Visittime int	`json:"visittime" form:"visittime" binding:"required"`
	UserName string		`json:"username" form:"username" binding:"required"`
	Telephone string		`json:"telephone" form:"telephone" binding:"required"`
	Email string	`json:"email" form:"email" binding:"required"`
}