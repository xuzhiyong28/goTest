package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// 虚拟节点个数
const DEFAULT_REPLICAS = 8

type HashRing []uint32

func (c HashRing) Len() int {
	return len(c)
}

func (c HashRing) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c HashRing) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type Node struct {
	Id       int
	Ip       string
	Port     int
	HostName string
	Weight   int
}

func NewNode(id int, ip string, port int, name string, weight int) *Node {
	return &Node{
		Id:       id,
		Ip:       ip,
		Port:     port,
		HostName: name,
		Weight:   weight,
	}
}

type Consistent struct {
	Nodes     map[uint32]Node
	numReps   int
	Resources map[int]bool
	ring      HashRing
	sync.RWMutex
}

func NewConsistent() *Consistent {
	return &Consistent{
		Nodes:     make(map[uint32]Node),
		numReps:   DEFAULT_REPLICAS,
		Resources: make(map[int]bool),
		ring:      HashRing{},
	}
}

// 添加节点
func (c *Consistent) Add(node *Node) bool {
	c.Lock()
	defer c.Unlock()
	// 判断node是否已经添加
	if _, ok := c.Resources[node.Id]; ok {
		return false
	}
	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		// c.hashStr(str) 会将str通过哈希函数转换成uint32数字
		c.Nodes[c.hashStr(str)] = *(node)
	}
	c.Resources[node.Id] = true
	c.sortHashRing()
	return true
}

func (c *Consistent) sortHashRing() {
	c.ring = HashRing{}
	for k := range c.Nodes {
		c.ring = append(c.ring, k)
	}
	sort.Sort(c.ring)
}

func (c *Consistent) joinStr(i int, node *Node) string {
	return node.Ip + "*" + strconv.Itoa(node.Weight) +
		"-" + strconv.Itoa(i) +
		"-" + strconv.Itoa(node.Id)
}
func (c *Consistent) hashStr(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) Get(key string) Node {
	c.RLock()
	defer c.RUnlock()

	hash := c.hashStr(key) //获取key的通过哈希函数计算的uint32数字
	i := c.search(hash)

	return c.Nodes[c.ring[i]]
}

func (c *Consistent) search(hash uint32) int {

	// 取hash的范围在哪一个范围
	i := sort.Search(len(c.ring), func(i int) bool {
		return c.ring[i] >= hash
	})
	if i < len(c.ring) {
		if i == len(c.ring)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(c.ring) - 1
	}
}

func (c *Consistent) Remove(node *Node) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.Resources[node.Id]; !ok {
		return
	}

	delete(c.Resources, node.Id)

	count := c.numReps * node.Weight
	for i := 0; i < count; i++ {
		str := c.joinStr(i, node)
		delete(c.Nodes, c.hashStr(str))
	}
	c.sortHashRing()
}

func main() {
	cHashRing := NewConsistent()
	for i := 0; i < 10; i++ {
		si := fmt.Sprintf("%d", i)
		cHashRing.Add(NewNode(i, "172.18.1."+si, 8080, "host_"+si, 1))
	}

	keys := make([]int, 0)
	for k, _ := range cHashRing.Nodes {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Hash:", k, " IP:", cHashRing.Nodes[uint32(k)].Ip)
	}
	/*for k, v := range cHashRing.Nodes {
		fmt.Println("Hash:", k, " IP:", v.Ip)
	}*/

	for i := 0; i < 20; i++ {
		si := fmt.Sprintf("key%d", i)
		k := cHashRing.Get(si)
		fmt.Printf("hash值 = %d 值 = %s, 节点IP = %s\n", cHashRing.hashStr(si), si, k.Ip)
	}

}
