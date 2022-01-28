package main

import (
	"encoding/hex"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/davecgh/go-spew/spew"
	"log"
	"net/http"
	"os"
	"text/template"
)

/**
	1. 先执行命令 rice embed-go  会生成rice-box.go文件
    2. 执行 go build 打包
    3. 这样以后可执行文件就能使用外部的文件了

*/

func main() {
	conf := rice.Config{
		LocateOrder: []rice.LocateMethod{rice.LocateEmbedded, rice.LocateAppended, rice.LocateFS},
	}
	box, err := conf.FindBox("example-files")
	if err != nil {
		log.Fatalf("error opening rice.Box: %s\n", err)
	}
	contentString, err := box.String("file.txt")

	if err != nil {
		log.Fatalf("could not read file contents as string: %s\n", err)
	}
	log.Printf("Read some file contents as string:\n%s\n", contentString)

	contentBytes, err := box.Bytes("file.txt")
	if err != nil {
		log.Fatalf("could not read file contents as byteSlice: %s\n", err)
	}
	log.Printf("Read some file contents as byteSlice:\n%s\n", hex.Dump(contentBytes))

	file, err := box.Open("file.txt")
	if err != nil {
		log.Fatalf("could not open file: %s\n", err)
	}
	spew.Dump(file)

	templateBox, err := rice.FindBox("example-templates")
	if err != nil {
		log.Fatal(err)
	}
	templateString, err := templateBox.String("message.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	// 解析和执行模板
	tmplMessage, err := template.New("message").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	tmplMessage.Execute(os.Stdout, map[string]string{"Message": "Hello, world!"})
	http.Handle("/", http.FileServer(box.HTTPBox()))
	go func() {
		fmt.Println("Serving files on :8080, press ctrl-C to exit")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("error serving files: %v", err)
		}
	}()
	select {}
}
