package controllers

import "MySeckill-BackEnd/models"

type UserController struct {
	BaseController
}

//获取用户列表
func (u *UserController) Get(){
	var student []models.Student
	student = models.GetStudentInfo()
	u.Data["student"] = student
	u.TplName = "user.html"
}

