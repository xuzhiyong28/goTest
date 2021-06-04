package pos

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"
)

// 区块结构
type Block struct {
	LastHash  string
	Hash      string
	TimeStamp string
	Data      string
	Height    int    //区块高度
	Address   string //出块的地址
}

//Block方法
func (b *Block) getHash() {
	sumString := b.LastHash + b.TimeStamp + b.Data + b.Address + strconv.Itoa(b.Height)
	hash := sha256.Sum256([]byte(sumString))
	b.Hash = hex.EncodeToString(hash[:])
}

//区块链
var BlockChain []Block

// 挖矿节点
type Node struct {
	tokens  int    //代币数量
	days    int    //币龄
	address string //节点地址
}

var mineNodePool []Node       //挖矿节点节点池
var probalityNodesPool []Node //获取代币的数量的概率
var randNodePool []Node       //随机节点池

// 初始化节点池
func init() {
	// 手动添加两个节点
	mineNodePool = append(mineNodePool, Node{1000, 1, "AAAAAAAAA"})
	mineNodePool = append(mineNodePool, Node{100, 3, "BBBBBBBBBBB"})
	// 初始化随机节点池, 挖矿概率与代币数量和币龄有关
	for _, v := range mineNodePool {
		for i := 0; i < v.tokens * v.days; i++ {
			randNodePool = append(randNodePool, v)
		}
	}
}

//每次挖矿都会从概率节点池中随机选出获得出块权的节点地址
func getMineAddress() string {
	bInt := big.NewInt(int64(len(randNodePool)))

	// 得出一个随机数,最大不超过随机节点池的大小
	rInt, err := rand.Int(rand.Reader, bInt)
	if err != nil {
		log.Panic(err)
	}
	return randNodePool[int(rInt.Int64())].address
}

//生成新区块
func generateNewBlock(oldBlock Block, data string, address string) Block {
	newBlock := Block{}
	newBlock.LastHash = oldBlock.Hash
	newBlock.Data = data
	newBlock.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	newBlock.Height = oldBlock.Height + 1
	newBlock.Address = getMineAddress()
	newBlock.getHash()
	return newBlock
}

func DoTestPos_1() {
	genesisBlock := Block{
		"0000000000000000000000000000000000000000000000000000000000000000",
		"",
		time.Now().Format("2006-01-02 15:04:05"),
		"我是创世区块", 1,
		"0000000000"}
	genesisBlock.getHash()
	//上链
	BlockChain = append(BlockChain, genesisBlock)
	fmt.Println(BlockChain[0])

	i := 0
	for {
		time.Sleep(time.Second)
		newBlock := generateNewBlock(BlockChain[i], "我是区块内容", "00000")
		BlockChain = append(BlockChain, newBlock)
		fmt.Println(BlockChain[i+1])
		i++
	}

}
