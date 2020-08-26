package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id                     int       `orm:"column(id);auto"`
	Username               string    `orm:"column(username);size(255)"`
	IsEmailVerified        int8      `orm:"column(is_email_verified)"`
	Salt                   string    `orm:"column(salt);size(32)"`
	PasswordHash           string    `orm:"column(password_hash);size(255)"`
	PasswordResetToken     string    `orm:"column(password_reset_token);size(255);null"`
	EmailConfirmationToken string    `orm:"column(email_confirmation_token);size(255);null"`
	Email                  string    `orm:"column(email);size(255)"`
	Role                   int16     `orm:"column(role)"`
	FromLdap               int16     `orm:"column(from_ldap)"`
	Status                 int16     `orm:"column(status)"`
	CreatedTime            time.Time `orm:"column(created_at);type(datetime)"`
	UpdatedTime            time.Time `orm:"column(updated_at);type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) InsertUser() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	return id, err
}

func (u *User) FindUserByNameExist(username string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(u)
	res := qs.Filter("username", username).Exist()
	return res
}

func (u *User) FindUserByName() error {
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("username", u.Username).One(u)
	return err
}

func (u *User) FindUserExistRaw() (bool, error) {
	o := orm.NewOrm()
	var username []string
	var email []string

	query := fmt.Sprintf("SELECT username,email FROM user where username='%s' or email='%s'", u.Username, u.Email)
	num, err := o.Raw(query).QueryRows(&username, &email)
	if err != nil {
		return true, err
	} else if num > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
