package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
)

type Ads struct {
	Id          int     `orm:"column(id);pk;auto" json:"-"`
	Name        string  `orm:"column(name)" json:"name"`
	Description string  `orm:"column(description)" json:"description,omitempty"`
	Img1        string  `orm:"column(img1)" json:"img1"`
	Img2        string  `orm:"column(img2);null" json:"img2,omitempty"`
	Img3        string  `orm:"column(img3);null" json:"img3,omitempty"`
	Price       float64 `orm:"column(price)" json:"price"`

	Base
}

func (t *Ads) TableName() string {
	return "ads"
}

func init() {
	orm.RegisterModel(new(Ads))
}

// AddAds insert a new Ads into database and returns
// last inserted Id on success.
func AddAds(m *Ads) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdsById retrieves Ads by Id. Returns error if
// Id doesn't exist
func GetAdsById(id int) (v *Ads, err error) {
	//o := orm.NewOrm()
	//v = &Ads{Id: id}
	//if err = o.Read(v); err == nil {
	//	return v, nil
	//}
	//return nil, err

	o := orm.NewOrm()
	v = &Ads{}
	err = o.QueryTable(new(Ads)).Filter("Id", id).One(v)

	return v, err
}

// GetAllAds retrieves all Ads matches certain condition. Returns empty list if
// no records exist
func GetAllAds(limit, offset int, sort, crease string) ([]*Ads, error) {
	o := orm.NewOrm()
	var ads []*Ads

	if sort == "date" {
		sort = "created_at"
	}

	if crease == "decrease" {
		sort = "-" + sort
	}

	_, err := o.QueryTable(new(Ads)).OrderBy(sort).Limit(limit).Offset(offset).All(&ads, "name", "img1", "price")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return ads, err
}
