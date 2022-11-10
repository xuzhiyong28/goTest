package main

import (
	"context"
	"example/plugin/go_kit/napodate"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

/**
http://localhost:8080/get
http://localhost:8080/status
curl -XPOST -d '{"date":"32/12/2020"}' http://localhost:8080/validate
*/
func main() {
	ctx := context.Background()
	srv := napodate.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	// 端点
	endpoints := napodate.Endpoints{
		GetEndpoint:      napodate.MakeGetEndpoint(srv),
		StatusEndpoint:   napodate.MakeStatusEndpoint(srv),
		ValidateEndpoint: napodate.MakeValidateEndpoint(srv),
	}
	go func() {
		handler := napodate.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(":8080", handler)
	}()
	<-errChan
}
