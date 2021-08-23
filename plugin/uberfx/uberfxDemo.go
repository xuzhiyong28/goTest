package uberfx

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
	"time"
)

func Demo1() {
	app := fx.New(
		fx.Provide(NewMyConstruct, NewHandler, NewMux, NewLogger),                            // 构造函数
		fx.Invoke(invokeNothingUse, invokeRegister, invokeAnotherFunc, invokeUseMyconstruct), // 初始化函数
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	http.Get("http://localhost:8080/")
	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}

func NewLogger(lc fx.Lifecycle) *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")

	lc.Append(fx.Hook{
		OnStart: func(i context.Context) error {
			logger.Println("logger onstart..")
			return nil
		},
		OnStop: func(i context.Context) error {
			logger.Println("logger onstop..")
			return nil
		},
	})
	return logger
}

// http.Handler 类型对象的构造函数，它输入参数中需要*log.Logger类型，所以在它执行之前，先执行NewLogger
func NewHandler(lc fx.Lifecycle, logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	lc.Append(fx.Hook{
		OnStart: func(i context.Context) error {
			logger.Println("handler onstart..")
			return nil
		},
		OnStop: func(i context.Context) error {
			logger.Println("handler onstop..")
			return nil
		},
	})

	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

// *http.ServeMux 类型对象的构造函数，它输入参数中需要*log.Logger类型，所以在它执行之前，先执行NewLogger
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

// 自定义类型
type mystruct struct {
	Name string
}

// mystruct 构造函数
func NewMyConstruct(logger *log.Logger) mystruct {
	logger.Println("Executing NewMyConstruct.")
	return mystruct{
		Name: "xuzhiyong",
	}
}

// invokeUseMyconstruct 是invoke函数，在其执行之前，需要执行 NewMyConstruct
func invokeUseMyconstruct(logger *log.Logger, c mystruct) {
	logger.Println("invokeUseMyconstruct..")
}

func invokeRegister(mux *http.ServeMux, h http.Handler, logger *log.Logger) {
	logger.Println("invokeRegiste...")
	mux.Handle("/", h)
}

func invokeAnotherFunc(logger *log.Logger) {
	logger.Println("invokeAnotherFunc...")
}

// 这个invoke函数，不依赖任何输入变量，所以不需要执行任何构造函数
func invokeNothingUse() {
	fmt.Println("invokeNothingUse...")
}
