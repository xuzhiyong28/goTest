package downloader

import (
	"fmt"
	"net/http"
	"testing"
)


func TestHead(t *testing.T) {
	resp, _ := http.Head("https://studygolang.com/dl/golang/go1.16.5.src.tar.gz")
	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		fmt.Println("bytes")
	}

}

func TestDownLoader(t *testing.T) {
	NewDownloader(1, true).Download("https://studygolang.com/dl/golang/go1.16.5.src.tar.gz", "go1.16.5.src.tar.gz")
}
