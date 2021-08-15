# Lottery
以太坊智能合约实战 —— FanOne彩票网

# 前言
> 这个项目是以前在网上学区块链的时候写的，现在才写出来做点笔记。
> 区块链小白一枚，也欢迎大家一起学习进步~

**业务逻辑分析**
- `remix`编写智能合约
    1. 函数：智能合约接收投注的以太币（`payable`）
    2. 状态变量：账户的保存（数组：[{0x00:1231242}, {0x01:1231242}, {0x00:1231243}]）
    3. 函数：查询账户信息（需要接收`账户地址`，`余额`，`投注信息`）
    4. 函数：开奖 (点击开奖按钮，随机生成中奖号码)
	    1.	 中奖号码状态变量: 存储中奖号码
	    2.  中奖用户状态变量：存储中奖用户

- `go-web`对接智能合约，渲染数据到前端
    1. 操作智能合约
	    1. 把投注人和对应的投注信息传到智能合约
	    2. 查询投注人的投注信息（投注的号码）
	    3. 查询指定账户的信息（余额）
	    4. 开奖：（查询开奖结果）

    2. 渲染数据到前端
	    1. 投注账户实现渲染到前端（长连接）
	    2. 渲染投注信息
	    3. 渲染余额信息

- 前端渲染
    1. 接收账户地址、账户交易密码，进行转账到智能合约
    2. 渲染信息
    3. 固定的智能合约代码

# 项目结构
```
lotteryWeb/
├── conf
├── controller
├── eth
│  ├── lottert——data
├── routers
├── sol
├── status
├── utils
└── views
```

conf:配置文件

controller: 编写业务逻辑

eth/lottery_data: 存放智能合约用户的信息等

eth: 编写与智能合约交互的接口

routers： 存放路由

sol : 存放合约

status : 存放静态文件

utils : 存放一些功能模块

views : 前端页面



大体效果展示

![请添加图片描述](https://img-blog.csdnimg.cn/0c4d633b7aab47828806bd48e1dc3179.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)


# 智能合约部分
生成智能合约对应的go文件
```shell script
abigen --abi lottery.abi --pkg eth --type Lottery --out lottery.go
```
启动私链
```shell script
geth --datadir "lottery_data" --identidy "fanone_lottery" --rpc --rpcport 8545 --rpcapi "eth,personal,web3,net,db" --rpccorsdomain "*" console
```
部署合约

![在这里插入图片描述](https://img-blog.csdnimg.cn/a752a045e5c24a02bd1634805ed6b2b0.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)


要通过挖矿来确定是否部署成功

![在这里插入图片描述](https://img-blog.csdnimg.cn/2047c75973064255b0a4e36551fbc68a.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)

这里出现这些信息就是成功了的

![在这里插入图片描述](https://img-blog.csdnimg.cn/7c0e482b4b8b411789a3960acea67f40.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NTMwNDUwMw==,size_16,color_FFFFFF,t_70)

# Go语言部分
下载依赖
```shell script
go mod tidy
```
运行即可
```shell script
go run main.go
```


> 这个项目是以前在网上学区块链的时候写的，现在才写出来做点笔记。
> 区块链小白一枚，也欢迎大家一起学习进步~

如果你喜欢的话，可以在右上角点一个`star`噢~ 

感谢阅读~