package routers

import (
	"lotteryWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.EthController{},"get:Index")		//首页
	beego.Router("/get-accounts", &controllers.EthController{},"get:GetBetAccounts") //获取投注账户
	beego.Router("/post-bet", &controllers.EthController{},"post:PostBet") 		     //进行投注
	beego.Router("/get-account-info", &controllers.EthController{},"get:GetAccountInfo")  //获取用户信息
	beego.Router("/search", &controllers.EthController{},"get:Search")  		  //查询
	beego.Router("/run-lottery", &controllers.EthController{},"get:RunLottery")    //开奖页面
	beego.Router("/run-lottery", &controllers.EthController{},"post:RunLottery")   //开奖
	beego.Router("/contract", &controllers.EthController{},"get:RunLottery")       //获取合约

}
