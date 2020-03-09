package main

import (
	_ "candy/routers"
	"github.com/astaxie/beego"
)

type MainController struct{
	beego.Controller
}
func (this *MainController)Get(){
	this.Ctx.WriteString("h")
}
func main() {
	beego.Router("/abc", &MainController{})
	beego.Run()
}
