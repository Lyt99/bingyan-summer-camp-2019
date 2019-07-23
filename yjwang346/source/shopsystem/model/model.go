package model

//用户登录以及注册时候用的
type Register struct{
	Username string		`json:"username" form:"username" binding:"required"`
	Password string		`json:"password" form:"password" binding:"required"`
	Nickname string		`json:"nickname" form:"nickname" binding:"required"`
	Mobile string		`json:"mobile" form:"mobile" binding:"required"`
	Email string	`json:"email" form:"email" binding:"required"`
}
type Login struct {
	Username string		`json:"username" form:"username" binding:"required"`
	Password string		`json:"password" form:"password" binding:"required"`
}
//用户信息页的内容，Visittime在什么时候要进行初始化
type Userinfo struct{
	//Id int		`json:"id" form:"id" binding:"required"`
	Username string		`json:"username" form:"username" binding:"required"`
	Password string		`json:"password" form:"password" `//binding:"required"
	Nickname string		`json:"nickname" form:"nickname" `//binding:"required"
	Mobile string		`json:"mobile" form:"mobile" `// binding:"required"
	Email string	`json:"email" form:"email" `//binding:"required"
	Visittime int	`json:"visittime" form:"visittime" `//binding:"required"
	Collcectcount int `json:"collcectcount" form:"collectcount"`
}

type Changeuserinfo struct {
	Password string		`json:"password" form:"password" binding:"required"`
	Nickname string		`json:"nickname" form:"nickname" `//binding:"required"
	Mobile string		`json:"mobile" form:"mobile" `// binding:"required"
	Email string	`json:"email" form:"email" `//binding:"required"
}

type Commodity struct {
	//Id用来表示序列号，还可以表示存入数据库的先后顺序
	Id int		`json:"id" form:"id" `//binding:"required"
	Pub_user string		`json:"pub_user" form:"pub_user" binding:"required"`
	Title string		`json:"title" form:"title" binding:"required"`
	Description string		`json:"desc" form:"desc" binding:"required"`
	Category int	`json:"category" form:"category" binding:"required"`
	Price float64		`json:"price" form:"price" binding:"required"`
	Picture string		`json:"picture" form:"picture" `//binding:"required"

	View_count int		`json:"view_count" form:"view_count" `//binding:"required"
	Collect_count int		`json:"collect_count" form:"collect_count" `//binding:"required"
}
type PostCommodity struct {
	//Id用来表示序列号，还可以表示存入数据库的先后顺序
	//Id int		`json:"id" form:"id" `//binding:"required"
	//Pub_user string		`json:"pub_user" form:"pub_user" binding:"required"`
	Title string		`json:"title" form:"title" binding:"required"`
	Description string		`json:"desc" form:"desc" binding:"required"`
	Category int	`json:"category" form:"category" binding:"required"`
	Price float64		`json:"price" form:"price" binding:"required"`
	Picture string		`json:"picture" form:"picture" `//binding:"required"
}