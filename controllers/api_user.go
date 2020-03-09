package controllers

type ApiUserController struct {
	ApiController
}
func (this *ApiUserController)Get(){
	this.Ctx.WriteString("user get")
}
func (this *ApiUserController)Transfer(){
	this.Ctx.WriteString("user transfer")
}
func (this *ApiUserController)Kick(){
	this.Ctx.WriteString("user kick")
}