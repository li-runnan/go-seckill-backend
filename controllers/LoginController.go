package controllers

import (
	"MySeckill-BackEnd/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

func (l *LoginController) Post() {
	username := l.GetString("username")
	password := md5.Sum([]byte(l.GetString("password")))
	md5password := hex.EncodeToString(password[:])

	admin := models.Admin{Username: username}
	o := orm.NewOrm()
	err := o.Read(&admin, "username")

	if err == orm.ErrNoRows || admin.Password != md5password {
		l.Ctx.ResponseWriter.Write([]byte("<script language='javascript'>alert('用户名或密码错误');window.location.href='/login';</script>"))
	} else {
		l.SetSession("username",username)
		l.SetSession("name",admin.Name)
		l.Redirect("/home",302)
	}
}
