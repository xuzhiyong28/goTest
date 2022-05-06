package hex

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// hex包实现了16进制字符表示的编解码
// https://itpika.com/2019/07/04/go/library-encoding-hex/

func TestDemo0(t *testing.T) {
	// 字符串转16进制
	byteData := []byte("测试数据")

	dataStr := hex.EncodeToString(byteData)
	fmt.Println(fmt.Sprintf("byte : %v , str : %v", byteData, dataStr))

	byteData2, _ := hex.DecodeString(dataStr)
	fmt.Println(fmt.Sprintf("byte : %v , str : %v", byteData2, dataStr))
}

func TestDemo1(t *testing.T) {
	srcCode := []byte("测试数据")
	dstEncode := make([]byte, hex.EncodedLen(len(srcCode)))
	hex.Encode(dstEncode, srcCode)
	fmt.Println(fmt.Sprintf("src : %v , dst : %v", srcCode, dstEncode))
}
