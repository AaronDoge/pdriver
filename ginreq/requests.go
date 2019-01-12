package ginreq

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

// 长连接
func ReqL() {
	client := &http.Client{
		Transport: &http.Transport{
			//Proxy: http.ProxyFromEnvironment,
			Proxy: http.ServeTLS,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        10000,
			MaxIdleConnsPerHost: 10000,
			IdleConnTimeout:     time.Duration(90) * time.Second,
		},
		Timeout: 20 * time.Second,
	}

	client.
}

func RequestUrl() {

	//生成client 参数为默认
	client := &http.Client{}

	//生成要访问的url
	url := "https://www.baidu.com"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)


	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println("req error", err.Error())
		return
	}

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)

}
