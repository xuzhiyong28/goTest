package downloader

import "testing"

func TestDownLoader(t *testing.T) {
	NewDownloader(10, true).Download("http://xxxx", "xxx.txt")
}
