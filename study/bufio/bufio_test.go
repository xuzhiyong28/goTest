package bufio

// https://zhuanlan.zhihu.com/p/73690883
// http://liuqh.icu/2021/04/13/go/package/6-bufio/
// bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，使用这个包可以大幅提高文件读写的效率
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDemo1(t *testing.T) {
	file, err := os.Open("D:\\temp\\Ftp And Frameio Comparison.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 10*1024)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}
