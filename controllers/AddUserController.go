package controllers

import (
	"MySeckill-BackEnd/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
)

type AddUserController struct {
	BaseController
}

func (a *AddUserController) Get(){
	a.TplName = "addUser.html"
}

func (a *AddUserController) Post(){
	studentId := a.GetString("username")
	password := md5.Sum([]byte("123456"))
	md5password := hex.EncodeToString(password[:])
	student := models.Student{Studentid: studentId,Password: md5password}
	o := orm.NewOrm()
	_, err := o.Insert(&student)

	if err != nil{
		panic(err)
	} else {
		a.Ctx.ResponseWriter.Write([]byte("<script language='javascript'>alert('添加成功');window.location.href='/addUser';</script>"))
	}
}
