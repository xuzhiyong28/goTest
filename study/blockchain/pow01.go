package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type block struct {
	//上一个区块的哈希
	PreHash string
	//当前区块的哈希
	HashCode string
	//时间戳
	TimeStamp string
	//难度系数
	Diff int
	//交易信息
	Data string
	//区块高度
	Index int
	//随机值
	Nonce int
}

//生成区块的哈希值
func GenerateFirstBlock(data string) block{
	var firstblock block
	firstblock.PreHash = "0"
	firstblock.TimeStamp = time.Now().String()
	//暂设为4
	firstblock.Diff = 4
	//交易信息
	firstblock.Data = data
	firstblock.Index = 1
	firstblock.Nonce = 0
	//通过sha256得到自己的哈希
	firstblock.HashCode = GenerationHashValue(firstblock)
	return firstblock
}

//生成区块的哈希值
func GenerationHashValue(block block) string{
	//按照比特币的写法，将区块的所有属性拼接后做哈希运算
	//int转为字符串
	var hashdata = strconv.Itoa(block.Index) + strconv.Itoa(block.Nonce) + strconv.Itoa(block.Diff) + block.TimeStamp
	//算哈希
	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

//产生新的区块
func GenerateNextBlock(data string, oldBolock block) block {
	//产生一个新的区块
	var newBlock block
	newBlock.TimeStamp = time.Now().String()
	//难度系数
	newBlock.Diff = 5
	//高度
	newBlock.Index = 2
	newBlock.Data = data
	newBlock.PreHash = oldBolock.HashCode
	newBlock.Nonce = 0

	//创建pow()算法的方法
	//计算前导0为4个的哈希值
	newBlock.HashCode = pow(newBlock.Diff, &newBlock)
	return newBlock
}

//pow算法
func pow(diff int , block *block) string{
	//不停的挖矿
	for {
		//认为是挖了一次矿了
		hash := GenerationHashValue(*block)
		//判断哈希值前导0是否为diff个0
		//strings.Repeat:判断hash是否有diff个0，写1，就判断为有多少个1
		if strings.HasPrefix(hash, strings.Repeat("0", diff)) {
			//挖矿成功
			fmt.Println("挖矿成功")
			return hash
		} else {
			//没挖到
			//随机值自增
			block.Nonce++
		}
	}
}

func DoTestPow_1(){
	//测试创建创世区块
	var firstBlock = GenerateFirstBlock("创世区块")
	fmt.Println(firstBlock)
	fmt.Println(firstBlock.Data)

	secondBlock := GenerateNextBlock("第二区块", firstBlock)
	fmt.Println(secondBlock)

	thirdBlock := GenerateNextBlock("第三区块", secondBlock)
	fmt.Println(thirdBlock)
}


