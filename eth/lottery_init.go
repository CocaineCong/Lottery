package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var LotteryObj *Lottery
var err error

func init() {
	client,err := ethclient.Dial("http://127.0.0.1:8545")  //连接以太坊智能合约
	if err != nil {
		fmt.Println("连接以太坊智能合约出错")
	}
	//合约地址
	contractAddr:="0xFAE1ed84956411e651C8B3abFcD21B131b3767f4"
	commonContractAddr := common.HexToAddress(contractAddr)
	LotteryObj ,err = NewLottery(commonContractAddr,client)  //实例化合约
	if err != nil {
		fmt.Println("实例化合约失败")
	}
}
