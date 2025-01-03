package http

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	POST = "POST"
	GET  = "GET"
)

var HttpClient *httpClient

func init() {
	HttpClient = &httpClient{}
}

type httpClient struct {
}

func (c *httpClient) Get(apiUrl string, params map[string]string) (string, error) {
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}

	var rsp *http.Response
	var err error

	rsp, err = http.Get(apiUrl + "?" + data.Encode())
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	rs, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	defer rsp.Body.Close()
	return string(rs), nil
}

func (c *httpClient) Post(apiUrl string, params map[string]string) (string, error) {
	data := url.Values{}
	for k, v := range params {
		data.Set(k, v)
	}
	var req *http.Request
	var rsp *http.Response
	var err error

	body := strings.NewReader(data.Encode())
	req, err = http.NewRequest(POST, apiUrl, body)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	cli := http.Client{}
	rsp, err = cli.Do(req)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	rs, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	defer rsp.Body.Close()
	return string(rs), nil
}

func (c *httpClient) PostJson(apiUrl string, jsonParams string) string {
	var jsonStr = []byte(jsonParams)
	body := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest(POST, apiUrl, body)
	if err != nil {
		log.Println("err:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	log.Println(apiUrl)
	log.Println(body)
	log.Println(req.Body)
	cli := http.Client{Timeout: 300 * time.Second}
	rsp, err := cli.Do(req)
	if err != nil {
		log.Println("do,err:", err)
		return ""
	}

	rs, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Println("readall,err:", err)
		return ""
	}
	defer rsp.Body.Close()
	return string(rs)
}
