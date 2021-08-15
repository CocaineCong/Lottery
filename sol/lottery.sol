pragma solidity 0.4.24;


contract Lottery {


    uint8 public total = 1;    // 彩票期数
    address public owner;      // 设置权限，只有管理员才能进行开奖

    constructor()public {
        owner = msg.sender;
    }


    modifier CheckOwner(){
        require(msg.sender == owner);    // 管理员才能进行开奖，需要认证
        _;
    }


    modifier CheckMoney() {
       require(msg.value == 1 ether);    // 校验转账的金额，必须是1 eth
       _;
    }


    struct Account {    // 投注者及投注信息结构体
        string addr;
        string nums;    // 1,2,3,4,2
    }


    Account[] public accounts;    // 所有投注者及投注信息
    Account public acount;        // 一个投注者及投注信息


    function Bet(string newAddr,string newNums) public payable CheckMoney{    // 投注函数
        acount = Account(newAddr,newNums);
        accounts.push(acount);
    }

    function GetAccountBalance(address addr) public view returns(uint256){
        return addr.balance;
    }

    uint256[] public retNums;       // 开奖结果
    mapping(uint8 =>  Account[]) public historyAccounts;    // 投注人和投注号码历史记录
    mapping(uint8 => uint256[]) public historyRetNums;      // 彩票期数，开奖结果历史记录

    function RunLottery() public CheckOwner {     // 开奖函数
        for(int i=1;i<=4;i++){                    // 红球从1-4里面随机选一个，选4次
            uint256 random1 = uint256(keccak256(abi.encodePacked(block.difficulty, int(now)+i)));
            uint256 num1 = random1 % 4 + 1;       // 去掉0的就+1
            retNums.push(num1);
        }

        // 蓝球从1-3里面随机选一个，选一次
        uint256 random2 = uint256(keccak256(abi.encodePacked(block.difficulty, int(now)+5)));
        uint256 num2 = random2 % 3 + 1;           // 去掉0的+1
        retNums.push(num2);

        historyAccounts[total] = accounts;        // 存储投注人和投注号码历史记录
        historyRetNums[total] = retNums;        // 存储彩票期数，开奖结果历史记录

        delete accounts;        // 清空投注记录
        delete retNums;         // 清空开奖结果数组

        total = total + 1;      //期数加一
    }
}