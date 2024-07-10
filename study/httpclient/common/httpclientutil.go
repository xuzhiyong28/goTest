package common

// https://segmentfault.com/a/1190000039086957 配置
import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadFile struct {
	// 表单名称
	Name string
	// 文件全路径
	Filepath string
}

type HttpClientUtil struct {
	client *http.Client
}

func NewHttpClientUtil(isProxy bool) *HttpClientUtil {
	if isProxy {
		proxy, _ := url.Parse("127.0.0.1:7890")
		tr := &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		return &HttpClientUtil{
			client: &http.Client{
				Transport: tr,
				Timeout:   time.Second * 5, //超时时间
			},
		}
	} else {
		return &HttpClientUtil{
			client: &http.Client{
				Timeout: time.Second * 5, //超时时间
			},
		}
	}
}

func (hc *HttpClientUtil) Get(reqUrl string, reqParams map[string]string, headers map[string]string) (string, error) {
	urlParams := url.Values{}
	url, _ := url.Parse(reqUrl)
	for key, val := range reqParams {
		urlParams.Set(key, val)
	}
	//进行URLEncode
	url.RawQuery = urlParams.Encode()
	//得到完整的url
	urlPath := url.String()
	httpRequest, _ := http.NewRequest("GET", urlPath, nil)
	// 添加请求头
	if headers != nil {
		for key, val := range headers {
			httpRequest.Header.Add(key, val)
		}
	}
	// 发起请求
	resp, err := hc.client.Do(httpRequest)
	if err != nil {
		return "", err
	}
	fmt.Println(resp.Header)
	fmt.Println(resp.Body)
	return "", nil
}
func (hc *HttpClientUtil) Post(reqUrl string, reqParams map[string]string, contentType string, files []UploadFile, headers map[string]string) (string, error) {
	requestBody, realContentType := hc.getReader(reqParams, contentType, files)
	httpRequest, _ := http.NewRequest("POST", reqUrl, requestBody)
	// 添加请求头
	httpRequest.Header.Add("Content-Type", realContentType)
	if headers != nil {
		for k, v := range headers {
			httpRequest.Header.Add(k, v)
		}
	}
	// 发送请求
	resp, err := hc.client.Do(httpRequest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(response)
	return "", nil
}

func (hc *HttpClientUtil) getReader(reqParams map[string]string, contentType string, files []UploadFile) (io.Reader, string) {
	if strings.Index(contentType, "json") > -1 {
		bytesData, _ := json.Marshal(reqParams)
		return bytes.NewReader(bytesData), contentType
	} else if files != nil {
		body := &bytes.Buffer{}
		// 文件写入 body
		writer := multipart.NewWriter(body)
		for _, uploadFile := range files {
			file, err := os.Open(uploadFile.Filepath)
			if err != nil {
				panic(err)
			}
			part, err := writer.CreateFormFile(uploadFile.Name, filepath.Base(uploadFile.Filepath))
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(part, file)
			file.Close()
		}
		// 其他参数列表写入 body
		for k, v := range reqParams {
			if err := writer.WriteField(k, v); err != nil {
				panic(err)
			}
		}
		if err := writer.Close(); err != nil {
			panic(err)
		}
		// 上传文件需要自己专用的contentType
		return body, writer.FormDataContentType()
	} else {
		urlValues := url.Values{}
		for key, val := range reqParams {
			urlValues.Set(key, val)
		}
		reqBody := urlValues.Encode()
		return strings.NewReader(reqBody), contentType
	}
}
