package uberfx

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"log"
	"os"
	"testing"
	"time"
)

type CustomizeStruct struct {
	Name string
}

func TestOne1(t *testing.T) {
	settings := Settings{
		modules: map[interface{}]fx.Option{},
		invokes: make([]fx.Option, 10),
	}

	op1 := Override(new(*log.Logger), func(lc fx.Lifecycle) *log.Logger {
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
	})

	opt2 := Override(new(CustomizeStruct), func(logger *log.Logger) CustomizeStruct {
		return CustomizeStruct{
			Name: "xuzhiyong",
		}
	})

	opt3 := Override(invoke(0), func(logger *log.Logger, c CustomizeStruct) {
		fmt.Println("=========== fx.Invoke init ==============")
	})

	if err := Options(op1, opt2, opt3)(&settings); err == nil {
		ctors := make([]fx.Option, 0, len(settings.modules))
		for _, opt := range settings.modules {
			ctors = append(ctors, opt)
		}

		stors := make([]fx.Option , 0 , len(settings.invokes))
		for _, opt := range settings.invokes {
			if opt != nil {
				stors = append(stors, opt)
			}
		}

		app := fx.New(fx.Options(ctors...), fx.Options(stors...))
		startCtx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
		defer cancel()
		if err := app.Start(startCtx); err != nil {
			log.Fatal(err)
		}
	}
}
