package controllers

type ApiLaunchGameController struct {
	ApiController
}
func (this *ApiLaunchGameController) Post(){
	this.Ctx.WriteString("launch game!!")
}