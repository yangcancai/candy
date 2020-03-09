package routers

import (
	"candy/controllers"
	"github.com/astaxie/beego"
)

// api 路由
func init(){
	// 获取游戏列表
	beego.Router("/api/gamelist", &controllers.ApiGameListController{})
	// 启动游戏
	beego.Router("/api/launchgame", &controllers.ApiLaunchGameController{})
	// 获取玩家额度
	beego.Router("/api/user/:user",&controllers.ApiUserController{})
	// 充值提现
	beego.Router("/api/user/:user/transfer",&controllers.ApiUserController{},"get,post:Transfer")
	// 玩家下线
	beego.Router("/api/user/:user/kick", &controllers.ApiUserController{}, "get,delete:Kick")
}
