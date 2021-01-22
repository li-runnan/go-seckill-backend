package routers

import (
	"MySeckill-BackEnd/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/home",&controllers.HomeController{})
    beego.Router("/welcome",&controllers.WelcomeController{})
    beego.Router("/logout",&controllers.LogoutController{})
    beego.Router("/user",&controllers.UserController{})
    beego.Router("/addUser",&controllers.AddUserController{})
    beego.Router("/ticket",&controllers.TicketController{})
    beego.Router("/publishTicket",&controllers.PublishTicketController{})
}
