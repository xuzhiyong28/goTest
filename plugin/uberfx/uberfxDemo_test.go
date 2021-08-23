package uberfx

import (
	"context"
	"go.uber.org/fx"
	"net/http"
	"testing"
)


func TestDemo1(t *testing.T) {
	Demo1()
}

func TestDemo2(t *testing.T) {
	register := func(lifecycle fx.Lifecycle) {
		mux := http.NewServeMux()
		server := http.Server{
			Addr: "8080",
			Handler: mux,
		}
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown(ctx)
			},
		})
	}

	fx.New(fx.Invoke(register)).Run()
}



