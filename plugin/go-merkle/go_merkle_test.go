package go_merkle

import (
	"crypto/sha256"
	"fmt"
	"github.com/xsleonard/go-merkle"
	"io/ioutil"
	"testing"
)

func splitData(data []byte, size int) [][]byte {
	count := len(data) / size
	blocks := make([][]byte, 0, count)
	for i := 0; i < count; i++ {
		block := data[i*size : (i+1)*size]
		blocks = append(blocks, block)
	}
	if len(data)%size != 0 {
		blocks = append(blocks, data[len(blocks)*size:])
	}
	return blocks
}

func TestSplitData(t *testing.T) {
	data, err := ioutil.ReadFile("testdata")
	if err != nil {
		fmt.Println(err)
		return
	}
	blocks := splitData(data, 32)
	tree := merkle.NewTree()
	err = tree.Generate(blocks, sha256.New())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Height: %d\n", tree.Height())
	fmt.Printf("Root: %v\n", tree.Root())
	fmt.Printf("N Leaves: %v\n", len(tree.Leaves()))
	fmt.Printf("Height 2: %v\n", tree.GetNodesAtHeight(2))

}
