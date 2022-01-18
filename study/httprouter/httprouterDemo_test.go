package httprouter

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"testing"
)

func TestDemo1(t *testing.T) {
	router := httprouter.New()

	// http://localhost:8080/
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Blog:%s \nWechat:%s", "www.flysnow.org", "flysnow_org")
	})

	// http://localhost:8080/user/flysnow
	router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

//一个新类型，用于存储域名对应的路由
type HostSwitch map[string]http.Handler

//实现http.Handler接口，进行不同域名的路由分发
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//根据域名获取对应的Handler路由，然后调用处理（分发机制）
	if hander := hs[r.Host]; hander != nil {
		hander.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func TestDemo2(t *testing.T) {
	playRouter := httprouter.New()
	playRouter.GET("/" , func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w,"hello, %s!\n", "playRouter")
	})
	toolRouter := httprouter.New()
	toolRouter.GET("/" , func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w,"hello, %s!\n", "toolRouter")
	})
	hs := make(HostSwitch)
	hs["play.flysnow.org:12345"] = playRouter
	hs["tool.flysnow.org:12345"] = toolRouter
	//HostSwitch实现了http.Handler,所以可以直接用
	log.Fatal(http.ListenAndServe(":12345", hs))
}
