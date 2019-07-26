package service

import (
	"final_task/model"
	"final_task/serializer"
)

type CommodityDeleteService struct {
	ID       int
	Username string
}

type CollectionDeleteService struct {
	ID       int
	Username string
}

func (c *CollectionDeleteService) DeleteCollection() *serializer.Response {
	collection := model.Collection{
		CommodityID: c.ID,
		Username:    c.Username,
	}
	s := model.GetMongoGlobalSession().Clone()
	defer s.Close()

	if err := s.DB(model.DBName).C(model.CollectionCollectName).Remove(collection); err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "删除收藏失败",
			Data:    nil,
		}
	}
	return nil

}

func (c *CommodityDeleteService) DeleteCommodity() *serializer.Response {
	commodity := model.Commodity{
		ID: c.ID,
	}
	if err := commodity.DeleteCommodityAsID(c.Username); err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "不存在商品",
			Data:    nil,
		}
	}
	if commodity.Desc != "delete" {
		return &serializer.Response{
			Success: false,
			Error:   "无法删除",
			Data:    nil,
		}
	}

	return &serializer.Response{
		Success: true,
		Error:   "",
		Data:    "ok",
	}
}
