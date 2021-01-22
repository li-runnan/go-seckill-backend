package controllers

import (
	"MySeckill-BackEnd/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type PublishTicketController struct {
	BaseController
}

func (p *PublishTicketController) Get(){
	p.TplName = "PublishTicket.html"
}

func (p *PublishTicketController) Post(){
	name := p.GetString("name")
	num, _ := p.GetInt("num")
	datetime1 := p.GetString("datetime1")
	startTime, err := time.Parse("2006-01-02 15:04:05",datetime1)
	if err != nil{
		panic(err)
	}
	fmt.Println(name,num,startTime)
	ticket := models.Tickets{Name: name,Num: num,Starttime: startTime}
	o := orm.NewOrm()
	_, err = o.Insert(&ticket)
	if err != nil{
		panic(err)
	} else{
		p.Ctx.ResponseWriter.Write([]byte("<script language='javascript'>alert('发布成功');window.location.href='/publishTicket';</script>"))
	}
}
