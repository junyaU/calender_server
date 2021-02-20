package controllers

import (
	"calender/models"
	"sort"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ScheduleController struct {
	beego.Controller
}

func (this *ScheduleController) RegisterSchedule() {
	id, _ := this.GetInt64("id")
	name := this.GetString("name")
	year := this.GetString("year")
	month := this.GetString("month")
	day := this.GetString("day")
	scheduledTime := this.GetString("scheduledTime")
	color := this.GetString("color")

	o := orm.NewOrm()
	user := models.User{Id: id}
	if err := o.Read(&user); err != nil {
		return
	}

	schedule := models.Schedule{}
	schedule.User = &user
	schedule.Name = name
	schedule.Year = year
	schedule.Month = month
	schedule.Day = day
	schedule.ScheduleTime = scheduledTime
	schedule.Color = color
	if _, err := o.Insert(&schedule); err != nil {
		return
	}

	this.Data["json"] = schedule
	this.ServeJSON()
}

func (this *ScheduleController) GetScheduleData() {
	userId := this.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	var scheduleArr []models.Schedule
	o.QueryTable(new(models.Schedule)).Filter("user_id", userId).All(&scheduleArr)

	sort.Slice(scheduleArr, func(i, j int) bool {
		layout := "15:04"
		timeVal, _ := time.Parse(layout, scheduleArr[i].ScheduleTime)
		nextTimeVal, _ := time.Parse(layout, scheduleArr[j].ScheduleTime)
		return timeVal.Before(nextTimeVal)
	})

	this.Data["json"] = scheduleArr
	this.ServeJSON()
}

func (this *ScheduleController) DeleteSchedule() {
	id, _ := this.GetInt64("id")
	o := orm.NewOrm()
	o.Delete(&models.Schedule{Id: id})
}

func (this *ScheduleController) EditSchedule() {
	id, _ := this.GetInt64("id")
	name := this.GetString("name")
	time := this.GetString("time")

	schedule := models.Schedule{Id: id}
	o := orm.NewOrm()
	if err := o.Read(&schedule); err != nil {
		return
	}

	schedule.Name = name
	schedule.ScheduleTime = time
	if _, err := o.Update(&schedule); err != nil {
		return
	}

	this.Data["json"] = schedule
	this.ServeJSON()
}
