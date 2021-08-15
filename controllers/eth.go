package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"lotteryWeb/eth"
	"math/big"
	"net/http"
	"strings"
	"time"
)

// 存放客户端的连接
var clients = make(map[*websocket.Conn]bool) //存放是否连接成功
//将http升级为ws
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
//设置通道
var accountsChan =make(chan map[string]interface{})

type EthController struct {
	beego.Controller
}


//首页
func (e *EthController) Index() {
	total,err := eth.GetTotal()
	if err != nil {
		fmt.Println(err)
	}
	e.Data["total"] = total
	e.TplName = "index.html"
}

//获取投注信息
func (e *EthController) PostBet() {
	username := e.GetString("username")
	password := e.GetString("password")
	first := e.GetString("one")
	second := e.GetString("second")
	third := e.GetString("third")
	fourth := e.GetString("fourth")
	fifth := e.GetString("fifth")

	nums := first + " "+second + " "+third + " "+fourth + " "+fifth
	//获取到前端传来的投注信息
	tx,isOk := eth.Bet(strings.ToLower(username),password,nums)
	if isOk {
		resData := map[string]interface{}{
			"code":200,
			"msg":"投注成功",
			"data":tx.Hash(),
		}
		e.Data["json"] = resData
		e.ServeJSON()
	}
	resData := map[string]interface{}{
		"code":400,
		"msg":"投注失败",
	}
	e.Data["json"] = resData
	e.ServeJSON()
}

//实时获取投注账户以投注号码，返回json
func (e *EthController) GetBetAccounts() {
	ws,err:=upgrader.Upgrade(e.Ctx.ResponseWriter,e.Ctx.Request,nil)
	if err != nil {
		fmt.Println(err)
	}
	//放到clients中
	clients[ws] = true
	fmt.Println("client",clients)
	for {
		//第一种方式
		time.Sleep(time.Second)
	}
}

//查询页面
func (e *EthController) Search() {
	e.TplName = "search.html"
}

// 获取查询数据
func (e *EthController) GetAccountInfo() {
	addr := e.GetString("addr")
	// 把addr传到智能合约，返回用户得账户余额和投注号码
	// 替换下面得balance,nums
	balance,err:=eth.GetAccountBalance(addr)
	if err != nil {
		balance = big.NewInt(0)
	}
	accounts := eth.GetBetAccounts()
	var nums []interface{}
	for _,accountMap := range accounts{
		addrRet := accountMap["addr"]
		if addrRet == addr {
			num := accountMap["nums"]
			nums = append(nums,num)
		}
	}
	retMap := map[string]interface{}{}
	if addr=="" {
		retMap = map[string]interface{}{
			"addr":addr,
			"balance":0,
			"nums":[][]int{},
		}
	}else{
		retMap = map[string]interface{}{
			"addr":addr,
			"balance":balance,
			"nums":nums,
		}
	}
	fmt.Println(retMap)
	e.Data["accountInfo"] =retMap
	e.TplName = "search.html"
}

//开奖页面
func (e *EthController) RunLotteryHtml() {
	total,err:=eth.GetTotal()
	if err != nil {
		total = 0
	}
	e.Data["total"] = total
	e.TplName = "run_lottery.html"
}

//开奖功能
func (e *EthController) RunLottery() {
	resultNum,total,err := eth.RunLottery()
	money,err:=eth.GetTotalMoney(total)
	time.Sleep(time.Second * 15)
	if err != nil {
		e.Data["ret_num"] = []*big.Int{}
		e.Data["admin"] = "30E18Dd5F8bC3197527D23d7269b9b6d52A4D677"
		e.Data["total"] = 0
		e.Data["money"] = 0
	} else {
		e.Data["ret_num"] = resultNum
		e.Data["admin"] = "30E18Dd5F8bC3197527D23d7269b9b6d52A4D677"
		e.Data["total"] = total
		e.Data["money"] = money
	}
	e.TplName = "run_lottery.html"
}

//智能合约页面
func (e *EthController) Contract()  {
	e.TplName = "contract.html"
}