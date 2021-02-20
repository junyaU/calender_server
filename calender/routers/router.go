// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"calender/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/createUser", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/emailCheck/?:email", &controllers.UserController{}, "get:EmailCheck")
	beego.Router("/registerSchedule", &controllers.ScheduleController{}, "post:RegisterSchedule")
	beego.Router("/getScheduleData/?:id", &controllers.ScheduleController{}, "get:GetScheduleData")
	beego.Router("/deleteSchedule", &controllers.ScheduleController{}, "post:DeleteSchedule")
	beego.Router("/editSchedule", &controllers.ScheduleController{}, "post:EditSchedule")
}
