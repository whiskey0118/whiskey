package websocket

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Device struct {
	Id           int       `orm:"column(id);auto"`
	Uid          string    `json:"uid" orm:"column(uid);size(255)"`
	Version      string    `json:"version" orm:"column(version);size(255)"`
	Area         string    `json:"area" orm:"column(area);size(255)"`
	Status       string    `json:"status" orm:"column(status);size(255)"`
	ExpiringDate time.Time `json:"expiring_date" orm:"column(expiring_date);type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Device))
}

func (d *Device) InsertUser() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(d)
	return id, err
}

func (d *Device) FindUserByName(uid string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(d)
	res := qs.Filter("username", uid).Exist()

	return res
}
