package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

var testApi = "http://api.proxy.ipidea.io/getProxyIp?num=100&return_type=json&lb=1&sb=0&flow=1&regions=&protocol=http"

func main() {
	getMyIp()

	//白名单代理
	var proxyIP = "ip:port"
	go httpProxy(proxyIP, "", "")
	go Socks5Proxy(proxyIP, "", "")
	time.Sleep(time.Minute)
}

func getMyIp() {
	rsp, err := http.Get("https://api.myip.la/en?json")
	if err != nil {
		fmt.Println("获取本机ip失败", err.Error())
		return
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), "本机ip:", string(body))

}

//http代理
func httpProxy(proxyUrl, user, pass string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), "http", "返回信息:", err)
		}
	}()
	urli := url.URL{}

	if !strings.Contains(proxyUrl, "http") {
		proxyUrl = fmt.Sprintf("http://%s", proxyUrl)
	}

	urlProxy, _ := urli.Parse(proxyUrl)
	if user != "" && pass != "" {
		urlProxy.User = url.UserPassword(user, pass)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlProxy),
		},
	}
	rqt, err := http.NewRequest("GET", testApi, nil)
	if err != nil {
		panic(err)
		return
	}
	response, err := client.Do(rqt)
	if err != nil {
		panic(err)
		return
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), proxyUrl, "【http success】", "返回信息:", response.Status, string(body))

	return
}

//socks5代理
func Socks5Proxy(proxyUrl, user, pass string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), proxyUrl, "返回信息:", err)
		}
	}()

	var userAuth proxy.Auth
	if user != "" && pass != "" {
		userAuth.User = user
		userAuth.Password = pass
	}
	dialer, err := proxy.SOCKS5("tcp", proxyUrl, &userAuth, proxy.Direct)
	if err != nil {
		panic(err)
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
				return dialer.Dial(network, addr)
			},
		},
		Timeout: time.Second * 10,
	}

	//请求网络
	if resp, err := httpClient.Get(testApi); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05 07"), proxyUrl, "【socks5 success】", "返回信息:", string(body))
	}
}
