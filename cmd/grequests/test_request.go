package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
	"net/url"
)

func main() {
	const reqUrl = "https://pknabi9jwilx.sw-cdnstream.com/hls2/01/00505/r3ycea9avtsh_,n,h,x,.urlset/master.m3u8?t=i5VByeF87LgnN9E9iLCs1sMIVIw-UeAYnWbz0dj1A3c&s=1694666519&e=129600&f=2525846&srv=jxyznjhecuim&i=0.4&sp=500&p1=jxyznjhecuim&p2=jxyznjhecuim&asn=199524"
	// 设置Clash代理的地址和端口
	proxyURL := "http://127.0.0.1:7890"

	// 创建HTTP代理URL
	proxyURLParsed, err := url.Parse(proxyURL)
	if err != nil {
		fmt.Println("无法解析代理URL:", err)
		return
	}

	// 创建一个自定义的http.Client，使用Clash代理
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		Proxy: http.ProxyURL(proxyURLParsed),
	//	},
	//}

	//// 创建HTTP请求
	//req, err := http.NewRequest("GET", reqUrl, nil)
	//if err != nil {
	//	fmt.Println("无法创建HTTP请求:", err)
	//	return
	//}
	//
	//// 发送HTTP请求
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("HTTP请求错误:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//// 检查响应状态码
	//if resp.StatusCode != 200 {
	//	fmt.Println("HTTP响应状态码:", resp.StatusCode)
	//	return
	//}
	//
	//// 处理响应数据
	//fmt.Println("HTTP响应内容:", resp.Status)
	//// 读取响应主体内容
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("无法读取响应主体:", err)
	//	return
	//}
	//
	//// 将响应主体内容转换为字符串并打印出来
	//fmt.Println("HTTP响应主体内容:")
	//fmt.Println(string(body))

	ro := &grequests.RequestOptions{
		Proxies: map[string]*url.URL{"https": proxyURLParsed, "http": proxyURLParsed},
	}
	resp, err := grequests.Get(reqUrl, ro)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())
	select {}
}
