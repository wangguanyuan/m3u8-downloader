package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// https://tl9clmrn.sw-cdnstream.com/hls2/01/02193/q4ya80l3to6u_,n,h,x,.urlset/master.m3u8?t=80CRIrgeOr3y7YFY4937bq5Yll3PtHMY4hbCxoDL8EY&s=1699412881&e=129600&f=10966788&srv=txmraegnjpog&i=0.4&sp=500&p1=txmraegnjpog&p2=txmraegnjpog&asn=16509
	// index-f2-v1-a1.m3u8?t=80CRIrgeOr3y7YFY4937bq5Yll3PtHMY4hbCxoDL8EY&s=1699412881&e=129600&f=10966788&srv=txmraegnjpog&i=0.4&sp=500&p1=txmraegnjpog&p2=txmraegnjpog&asn=16509
	// https://tl9clmrn.sw-cdnstream.com/hls2/01/02193/q4ya80l3to6u_,n,h,x,.urlset/index-f2-v1-a1.m3u8?t=80CRIrgeOr3y7YFY4937bq5Yll3PtHMY4hbCxoDL8EY&s=1699412881&e=129600&f=10966788&srv=txmraegnjpog&i=0.4&sp=500&p1=txmraegnjpog&p2=txmraegnjpog&asn=16509
	// index-f3-v1-a1.m3u8?t=80CRIrgeOr3y7YFY4937bq5Yll3PtHMY4hbCxoDL8EY&s=1699412881&e=129600&f=10966788&srv=txmraegnjpog&i=0.4&sp=500&p1=txmraegnjpog&p2=txmraegnjpog&asn=16509
	// https://tl9clmrn.sw-cdnstream.com/hls2/01/02193/q4ya80l3to6u_,n,h,x,.urlset/index-f3-v1-a1.m3u8?t=80CRIrgeOr3y7YFY4937bq5Yll3PtHMY4hbCxoDL8EY&s=1699412881&e=129600&f=10966788&srv=txmraegnjpog&i=0.4&sp=500&p1=txmraegnjpog&p2=txmraegnjpog&asn=16509
	const reqUrl = "https://dulisv62.sw-cdnstream.com/hls2/01/02109/0v2aao2ivdz2_,n,h,x,.urlset/index-f3-v1-a1.m3u8?t=yHazDyDTkfV0oxFZefXRKOLyBwSQYJcDtkjUg9ir97Q&s=1699260980&e=129600&f=10547473&srv=lrphxnueqxzy&i=0.4&sp=500&p1=lrphxnueqxzy&p2=lrphxnueqxzy&asn=4134"
	// 设置Clash代理的地址和端口
	proxyURL := "http://127.0.0.1:7890"

	// 创建HTTP代理URL
	proxyURLParsed, err := url.Parse(proxyURL)
	if err != nil {
		fmt.Println("无法解析代理URL:", err)
		return
	}

	// 创建一个自定义的http.Client，使用Clash代理
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURLParsed),
		},
	}

	// 创建HTTP请求
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		fmt.Println("无法创建HTTP请求:", err)
		return
	}
	req.Header.Set("Host", "g7sun940ri.sw-cdnstream.com")
	req.Header.Set("Origin", "https://sfastwish.com")
	req.Header.Set("Referer", "https://sfastwish.com/")

	// 发送HTTP请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP请求错误:", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != 200 {
		fmt.Println("HTTP响应状态码:", resp.StatusCode)
		return
	}

	// 处理响应数据
	fmt.Println("HTTP响应内容:", resp.Status)
	// 读取响应主体内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("无法读取响应主体:", err)
		return
	}

	// 将响应主体内容转换为字符串并打印出来
	fmt.Println("HTTP响应主体内容:")
	fmt.Println(string(body))
}
