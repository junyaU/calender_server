package models

import "time"

type Schedule struct {
	Id           int64     `orm:"auto"`
	User         *User     `orm:"rel(fk)"`
	Name         string    `orm:"size(20)"`
	Year         string    `orm:"size(10)"`
	Month        string    `orm:"size(10)"`
	Day          string    `orm:"size(10)"`
	ScheduleTime string    `orm:"size(30)"`
	Color        string    `orm:"size(20)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}
