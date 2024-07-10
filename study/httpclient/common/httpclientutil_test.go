package common

import "testing"

func TestNewHttpClientUtil(t *testing.T) {
	httpclientUtil := NewHttpClientUtil(true)
	httpclientUtil.Get("", nil, nil)
}
