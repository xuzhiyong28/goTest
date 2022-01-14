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
