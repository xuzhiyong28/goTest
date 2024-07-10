package contextsty

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestContextWithCancelDemo(t *testing.T) {
	ContextWithCancelDemo()
}

func TestContextWithTimeOutDemo(t *testing.T) {
	ContextWithTimeOutDemo()
}

func TestWaitGroupDemo(t *testing.T) {
	WaitGroupDemo0()
	WaitGroupDemo1()
}

func TestOther(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	go func(cancel context.CancelFunc) {
		time.Sleep(1 * time.Second)
		cancel()
	}(cancelFunc)
	select {
	case <-ctx.Done():
		fmt.Sprint("done")
	}
}

func TestServer(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number := rand.Intn(2)
		if number == 0 {
			time.Sleep(time.Second * 10) // 耗时10秒的慢响应
			fmt.Fprintf(w, "slow response")
			return
		}
		fmt.Fprint(w, "quick response")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type respData struct {
	resp *http.Response
	err  error
}

func TestTimeOutClient(t *testing.T) {
	// 定义一个100毫秒的超时
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancelFunc()
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *respData, 1)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/", nil)
	if err != nil {
		fmt.Printf("new requestg failed, err:%v\n", err)
		return
	}
	req = req.WithContext(ctx) // 使用带超时的ctx创建一个新的client request
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}
