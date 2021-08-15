package main

import (
	"github.com/astaxie/beego/toolbox"
	"lotteryWeb/controllers"
	_ "lotteryWeb/routers"
	"github.com/astaxie/beego"
)

func main() {
	controllers.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()
	beego.Run()
}

