package model



import "gopkg.in/mgo.v2/bson"

type Collection struct {
	Username    string `json:"username" bson:"username"`
	CommodityID int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
}

type CollectionBreviary struct {
	CommodityID int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
}

type CollectionBreviaryList struct {
	CBL []CollectionBreviary
}

func (c *Collection) RegisterCollection() error {
	s := GetMongoGlobalSession().Clone()
	defer s.Close()

	return s.DB(DBName).C(CollectionCollectName).Insert(c)
}

func (cBL *CollectionBreviaryList) GetCollections(username string) error {
	s := GetMongoGlobalSession().Clone()
	defer s.Close()

	iter := s.DB(DBName).C(CollectionCollectName).Find(bson.M{
		"username": username,
	}).Iter()
	var c Collection
	var cB CollectionBreviary
	for iter.Next(&c) {
		cB.CommodityID = c.CommodityID
		cB.Title = c.Title
		cBL.CBL = append(cBL.CBL, cB)
	}
	return iter.Close()
}
