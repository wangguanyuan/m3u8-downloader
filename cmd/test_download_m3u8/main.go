package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/grafov/m3u8"
)

func main() {
	// M3U8文件的URL
	m3u8URL := "https://example.com/path/to/video.m3u8"

	// 保存视频的文件名
	outputFile := "output.ts"

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

	// 发送HTTP请求获取M3U8文件
	resp, err := client.Get(m3u8URL)
	if err != nil {
		fmt.Println("无法获取M3U8文件:", err)
		return
	}
	defer resp.Body.Close()

	// 解析M3U8文件
	playlist, listType, err := m3u8.DecodeFrom(resp.Body, true)
	if err != nil {
		fmt.Println("无法解析M3U8文件:", err)
		return
	}

	// 确保M3U8文件类型是主播放列表
	if listType != m3u8.MEDIA {
		fmt.Println("不是主M3U8播放列表")
		return
	}

	// 创建输出文件
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer output.Close()

	// 遍历播放列表中的片段并下载
	segments := playlist.(*m3u8.MediaPlaylist).Segments
	for _, segment := range segments {
		// 构建片段的完整URL
		segmentURL := segment.URI
		if !strings.HasPrefix(segmentURL, "http") {
			segmentURL = m3u8URL[:strings.LastIndex(m3u8URL, "/")+1] + segmentURI
		}

		// 发送HTTP请求获取片段
		resp, err := client.Get(segmentURL)
		if err != nil {
			fmt.Println("无法获取片段:", err)
			return
		}
		defer resp.Body.Close()

		// 将片段写入输出文件
		_, err = io.Copy(output, resp.Body)
		if err != nil {
			fmt.Println("无法写入输出文件:", err)
			return
		}
	}

	fmt.Println("M3U8视频下载完成并保存为:", outputFile)
}
