package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Demo1() {
	router := gin.Default()
	router.GET("/someGet", func(context *gin.Context) {
		name := context.Param("name")
		fmt.Printf("Hello %s\n", name)
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
