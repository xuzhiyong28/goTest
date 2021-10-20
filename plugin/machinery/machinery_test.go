package machinery

import "github.com/RichardKnop/machinery/v1/config"

/**
更多实例可以看 D:\webspack\goCode\pkg\mod\github.com\!richard!knop\machinery@v1.7.4\example\machinery.go
*/

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"testing"
	"time"
)

const configPath = "config.yml"

func loadConfig() (*config.Config, error) {
	if configPath != "" {
		return config.NewFromYaml(configPath, true)
	}
	return config.NewFromEnvironment(true)
}

func startServer() (*machinery.Server, error) {
	cnf, err := loadConfig()
	if err != nil {
		return nil, err
	}
	server, err := machinery.NewServer(cnf) // 实例化
	if err != nil {
		return nil, err
	}
	tasks := map[string]interface{}{
		"add": func(args ...int64) (int64, error) {
			sum := int64(0)
			for _, arg := range args {
				sum += arg
			}
			return sum, nil
		},
		"multiply": func(args ...int64) (int64, error) {
			sum := int64(1)
			for _, arg := range args {
				sum *= arg
			}
			return sum, nil
		},
		"sum_ints": func(numbers []int64) (int64, error) {
			var sum int64
			for _, num := range numbers {
				sum += num
			}
			return sum, nil
		},
		"sum_floats": func(numbers []float64) (float64, error) {
			var sum float64
			for _, num := range numbers {
				sum += num
			}
			return sum, nil
		},
		"concat": func(strs []string) (string, error) {
			var res string
			for _, s := range strs {
				res += s
			}
			return res, nil
		},
	}
	return server, server.RegisterTasks(tasks)
}

func TestTask_Demo1(t *testing.T) {
	server, err := startServer()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 启动一个worker处理
	worker := server.NewWorker("asong", 1)
	go func() {
		err := worker.Launch()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("======= task - complete =======")
	}()

	// 发送任务  - 1
	asyncResult01, err := server.SendTaskWithContext(context.Background(), &tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	})
	results, err := asyncResult01.Get(time.Millisecond * 5)
	if err != nil {
		fmt.Errorf("Getting task result failed with error: %s", err.Error())
	}
	fmt.Printf("1 + 1 = %v\n", tasks.HumanReadableResults(results))
}

// 延迟队列
func TestTaskDelay_Demo(t *testing.T) {
	server, err := startServer()
	if err != nil {
		fmt.Println(err.Error())
	}
	//启动一个工作队列
	worker := server.NewWorker("asong", 1)
	go func() {
		err := worker.Launch()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("======= task - complete =======")
	}()
	eta := time.Now().Add(time.Millisecond * 5)
	// 指定一个延迟任务
	signature :=&tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 20,
			},
			{
				Type:  "int64",
				Value: 30,
			},
		},
		RetryCount: 3, //重试次数
		ETA: &eta,	//延迟秒数
	}
	_ , err = server.SendTask(signature)
	if err != nil {
		fmt.Println(err.Error())
	}
	time.Sleep(5 * time.Minute)
}
