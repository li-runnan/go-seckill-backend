package controllers

import "MySeckill-BackEnd/models"

type TicketController struct {
	BaseController
}

func (t *TicketController) Get(){
	var ticket []models.Tickets
	ticket = models.GetTicketsInfo()
	t.Data["ticket"] = ticket
	t.TplName = "ticket.html"
}
