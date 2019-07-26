package service

import (
	"final_task/model"
	"final_task/serializer"
	"gopkg.in/mgo.v2/bson"
)

type CommodityListGetService struct {
	Page     int    `json:"page" bson:"page" form:"page"`
	Limit    int    `json:"limit" bson:"limit" form:"limit"`
	Category int    `json:"category" bson:"category" form:"category"`
	Keyword  string `json:"keyword" bson:"keyword" form:"keyword"`
}

type HotKeywordGetService struct {
}

type MyCommodityListService struct {
	Username string
}

type MyCollectionsGetService struct {
	Username string `json:"username" bson:"username" form:"username"`
}

func (m *MyCollectionsGetService) GetMyCollections() (err *serializer.Response, mBL *model.CollectionBreviaryList) {
	s := model.GetMongoGlobalSession().Clone()
	defer s.Close()

	s.DB(model.DBName).C(model.CollectionCollectName).Find(bson.M{
		"username": m.Username,
	})

	if err, mBL = m.GetMyCollections(); err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "迭代器关闭出错",
			Data:    nil,
		}, nil
	}

	return nil, mBL
}

func (m *MyCommodityListService) GetMyCommodityList() *serializer.Response {
	var cBL model.CommodityBreviaryList
	if err := cBL.GetCommodityAsUserID(m.Username); err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "没有发布过商品或者查询出现了错误",
			Data:    "",
		}
	}
	return &serializer.Response{
		Success: true,
		Error:   "",
		Data:    nil,
	}
}

func (h *HotKeywordGetService) GetHotKeyword() *serializer.Response {
	s := model.GetMongoGlobalSession().Clone()
	defer s.Close()
	q := s.DB(model.DBName).C(model.KeywordCollectionName).Find(nil)
	if i, err := q.Count(); i == 0 || err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "查询关键词出错",
			Data:    nil,
		}
	} else if i == 1 {
		var data string
		q.Iter().Next(&data)
		return &serializer.Response{
			Success: true,
			Error:   "",
			Data:    data,
		}
	} else {
		var data [2]string
		iter := q.Sort().Iter()
		iter.Next(data[0])
		iter.Next(data[1])
		return &serializer.Response{
			Success: true,
			Error:   "",
			Data:    data,
		}
	}

}

func (c *CommodityListGetService) GetAllCommodityList() *serializer.Response {
	s := model.GetMongoGlobalSession().Clone()
	defer s.Close()

	if err := s.DB(model.DBName).C(model.KeywordCollectionName).Insert(bson.M{"keyword": c.Keyword}); err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "创建关键词失败",
			Data:    nil,
		}
	}
	commodityList := model.CommodityList{}
	if err := commodityList.FindAllCommodity(c.Category, c.Page, c.Limit, c.Keyword); err == nil {
		return &serializer.Response{
			Success: false,
			Error:   "没有找到",
			Data:    nil,
		}
	}
	a := serializer.CommoditiesListResponseArray{}
	a.FindAllSuccessResponse(commodityList.CList)
	return &serializer.Response{
		Success: true,
		Error:   "",
		Data:    a.CList,
	}
}
