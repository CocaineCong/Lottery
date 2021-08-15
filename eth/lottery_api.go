package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"io/ioutil"
	"lotteryWeb/utils"
	"math/big"
	"strings"
	"time"
)


// geth --datadir "lottery_data" --identity "fanone_lottery" --rpc --rpcport 8545 --rpcapi "eth,personal.web3,net,db" --rpccorsdomain "*" console

//获取管理员账户地址
func GetOwner() (owner common.Address,err error) {
	owner,err = LotteryObj.Owner(nil)
	if err != nil {
		return owner,err
	}
	return owner,nil
}

//获取彩票期数
func GetTotal() (total uint8, err error) {
	total ,err = LotteryObj.Total(nil)
	if err != nil {
		return total,err
	}
	return total,nil
}

//投注
func Bet(username string, password string, nums string) (tx *types.Transaction, isOk bool) {
	dir := `D:\CodeProjects\GoLandProjects\BlockChain\lottery\lotteryWeb\eth\lottery_data\keystore\`
	filename,exist := utils.GetFileName(dir,username)
	if exist {
		var keystorePath string = dir + filename
		byteData,errFile := ioutil.ReadFile(keystorePath)
		if errFile != nil {
			fmt.Println("errFile",errFile)
			panic("读取key失败")
		}
		//初始化交易信息
		auth,errAuth := bind.NewTransactorWithChainID(strings.NewReader(string(byteData)),password,big.NewInt(15))
		if errAuth != nil {
			panic("初始化交易出错")
		}

		auth.GasLimit = uint64(3000000)
		auth.GasPrice = big.NewInt(20000)
		auth.Value = big.NewInt(1000000000000000000)

		tx ,err := LotteryObj.Bet(auth,username,nums)
		fmt.Println("auth",auth)
		fmt.Println("username",username)
		fmt.Println("nums",nums)
		if err != nil {
			return tx, false
		}
		return tx,true
	}else{
		return tx,false
	}
}

//获取投注信息
func GetBetAccounts() []map[string]interface{} {
	var addrsMap []map[string]interface{}
	var i int64 = 0
	for {
		addrStruct,err := LotteryObj.Accounts(nil,big.NewInt(i))
		if err != nil {
			break
		}
		addr := addrStruct.Addr
		nums := addrStruct.Nums
		newNums := strings.Split(nums," ")
		addrMap := map[string]interface{}{}
		addrMap["addr"] = addr
		addrMap["nums"] = newNums
		addrsMap = append(addrsMap,addrMap)
		i++
	}
	return addrsMap
}

//获取指定账户得余额
func GetAccountBalance(addr string) (*big.Int, error) {
	balance,err:=LotteryObj.GetAccountBalance(nil,common.HexToAddress(addr))
	return balance,err
}

//开奖结果，一个是开奖号码，一个是奖励金额
func RunLottery() ([]*big.Int, uint8, error) {
	total,_:=GetTotal()
	keyStoryPath := `D:\CodeProjects\GoLandProjects\BlockChain\lottery\lotteryWeb\eth\lottery_data\keystore\UTC--2021-08-08T02-07-37.052523300Z--30e18dd5f8bc3197527d23d7269b9b6d52a4d677`
	byteData,errFile := ioutil.ReadFile(keyStoryPath)
	if errFile != nil {
		panic("读取key出错")
	}
	password := "12345678"
	//初始化交易信息
	auth, errAuth := bind.NewTransactorWithChainID(strings.NewReader(string(byteData)),password,big.NewInt(15))
	if errAuth!=nil {
		fmt.Println("errAuth",errAuth)
		panic("初始化交易出错")
	}
	_,err := LotteryObj.RunLottery(auth)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 8)
	var resultNum []*big.Int
	//fmt.Println(err)
	//fmt.Println(total)
	if err == nil { //如果没有发生错误
		for i := 0; i < 5; i++ {
			num,err := LotteryObj.HistoryRetNums(nil,total,big.NewInt(int64(i)))
			if err != nil {
				return resultNum,total,err
			}
			resultNum = append(resultNum,num)
		}
		return resultNum,total,nil
	}
	return resultNum,total,err
}

// 获取当前奖金金额
func GetTotalMoney(total uint8) (*big.Int, error) {
	//money,err:=LotteryObj.HistoryRetNums(nil,total)
	//time.Sleep(time.Second*3)
	return nil,err
}



