package model

//用户登录以及注册时候用的
type Register struct{
	Username string		`json:"username" form:"username" binding:"required"`
	Password string		`json:"password" form:"password" binding:"required"`
	Nickname string		`json:"nickname" form:"nickname" binding:"required"`
	Mobile string		`json:"mobile" form:"mobile" binding:"required"`
	Email string	`json:"email" form:"email" binding:"required"`
}

//用户信息页的内容，Visittime在什么时候要进行初始化
type Userinfo struct{
	//Id int		`json:"id" form:"id" binding:"required"`
	Username string		`json:"username" form:"username" binding:"required"`
	//Password string		`json:"password" form:"password" binding:"required"`
	Visittime int	`json:"visittime" form:"visittime" `//binding:"required"
	Nickname string		`json:"nickname" form:"nickname" `//binding:"required"
	Mobile string		`json:"mobile" form:"mobile" `// binding:"required"
	Email string	`json:"email" form:"email" `//binding:"required"
	Collcectcount int `json:"collcectcount" form:"collectcount"`
}