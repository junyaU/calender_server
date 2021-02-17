package models

import (
	"time"
)

type User struct {
	Id       int64     `orm:"auto"`
	Name     string    `orm:"size(40)"`
	Email    string    `orm:"size(100)"`
	Password string    `orm:"size(255)"`
	JwtToken string    `orm:"null;size(255)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}
