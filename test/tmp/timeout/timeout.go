package main

import (
	"net/http"
	"fmt"
	"io"
	"context"
	"time"
	//"io/ioutil"
)

func main()  {
	var (
		//resp *http.Response
		//content []byte
		err error
	)
	if _, err = client(); err != nil {
		fmt.Println(err.Error())
	}

	if _, err = randTrip(); err != nil {
		fmt.Println(err.Error())
	}
	//if content, err = ioutil.ReadAll(resp.Body); err != nil {
	//	fmt.Errorf(err.Error())
	//}
	//fmt.Println(string(content))
	//fmt.Println(err)
}

func client() (*http.Response, error) {
	var (
		c *http.Client
	)
	c = http.DefaultClient
	c.Timeout = time.Millisecond*1
	return c.Get("http://www.baidu.com")
}

func randTrip() (*http.Response, error) {
	var (
		err error
		r io.Reader
		req *http.Request
		ctx context.Context
		cancel context.CancelFunc

	)

	if req, err = http.NewRequest("GET", "http://www.baidu.com", r); err != nil {
		fmt.Errorf(err.Error())
	}
	ctx = req.Context()
	ctx, cancel = context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()
	req = req.WithContext(ctx)
	transport := http.DefaultTransport
	return transport.RoundTrip(req)
}