package controllers

type ApiGameListController struct {
	ApiController
}
func (this *ApiGameListController) Get() {
	//this.Ctx.WriteString("hello")
}
// controller之外扩展的prepare接口函数
func (this *ApiGameListController) ExtPrepare(){
}
