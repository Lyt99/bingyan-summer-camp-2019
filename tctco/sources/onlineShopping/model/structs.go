package model

type User struct {
	ID                int    `json:"-"`
	Username          string `json:"username"`
	Password          string `json:"-"`
	Nickname          string `json:"nickname"`
	Mobile            string `json:"mobile"`
	Email             string `json:"email"`
	TotalViewCount    int    `json:"total_view_count, omitempty"`
	TotalCollectCount int    `json:"total_collect_count, omitempty"`
}

type Commodity struct {
	ID           int     `json:",omitempty"`//no space between , and omitempty!
	Title        string  `json:"title"`
	Desc         string  `json:"desc"`
	Category     int     `json:"category"`
	Price        float64 `json:"price"`
	Picture      string  `json:"picture"`
	PubUser      int     `json:"pub_user"`
	ViewCount    int     `json:"view_count"`
	CollectCount int     `json:"collect_count"`
}

type CommodityBrief struct {
	ID       int
	Title    string
	Price    float64
	Category int
	Picture  string
}

type RegisterJSON struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserUpdateJSON struct {
	Password string `json:"password"`
	Nickname string `json:"nickname" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type ProductJSON struct {
	Title    string  `json:"title" binding:"required"`
	Desc     string  `json:"desc" binding:"required"`
	Category int     `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Picture  string  `json:"picture" binding:"required"`
	PubUser  int
}

type SearchJSON struct {
	PageNo   int    `json:"page" binding: "required"`
	PageSize int    `json:"limit" binding: "required"`
	Category int    `json:"category" binding: "required"`
	Keyword  string `json: "keyword" binding: "required"`
}

type MyCommodity struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
}

type MyCollection MyCommodity

type CollectionJSON struct {
	CommodityID string `json:"id" binding: "required"`
}
