package controllers

import (
	"github.com/astaxie/beego/toolbox"
	"lotteryWeb/eth"
)

func InitTask() {
	task := toolbox.NewTask("GetAccountAndWtChan","0/3 * * * * *",GetAccountsAndWtChan)
	toolbox.AddTask("GetAccountsAndWtChan",task)
}

//读取投注信息并写入通道
func GetAccountsAndWtChan() error {
	// 获取投注信息
	accounts := eth.GetBetAccounts()
	result := map[string]interface{}{
		"account":accounts,
	}
	accountsChan <- result
	return nil
}