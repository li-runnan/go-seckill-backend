package controllers

type HomeController struct {
	BaseController
}

func (h *HomeController) Get(){
	h.TplName = "index.html"
}
