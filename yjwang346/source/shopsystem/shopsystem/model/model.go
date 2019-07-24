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
	Id int		`json:"id" form:"id" `//binding:"required"
	//Pub_user string		`json:"pub_user" form:"pub_user" binding:"required"`
	Title string		`json:"title" form:"title" binding:"required"`
	Description string		`json:"desc" form:"desc" binding:"required"`
	//由于desc似乎有特殊的用途，所以命名部分采用description但是传递仍是按照API要求的使用desc
	Category int	`json:"category" form:"category" binding:"required"`
	Price float64		`json:"price" form:"price" binding:"required"`
	Picture string		`json:"picture" form:"picture" `//binding:"required"
}

type HotSearch struct {
	Category int	`json:"category" form:"category" `//binding:"required"
	Keyword string		`json:"keyword" form:"keyword" `//binding:"required"
	//是不是都不是必须要求的呢？再思考
}

//获得商品列表时候的request
type Findcommodity struct {
	Page int `json:"page" form:"page" binding:"required"`
	Limit int `json:"limit" form:"limit" binding:"required"`
	Category int `json:"category" form:"category" binding:"required"`
	Keyword string `json:"keyword" form:"keyword" `//binding:"required"
}