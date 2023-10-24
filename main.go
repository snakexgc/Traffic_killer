package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func downloadFile(url, filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to download %s: %v\n", url, err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Failed to save file %s: %v\n", filename, err)
		return
	}

	// 删除文件
	err = os.Remove(filename)
	if err != nil {
		fmt.Printf("Failed to delete file %s: %v\n", filename, err)
	}
}

func main() {
	urls := "https://speed.darknetwork.online"
	fmt.Printf("流量杀手已开始运行！请耐心等待流量消耗！")
	totalIterationsFlag := flag.Int("r", 1, "总消耗流量数（GB）")
	flag.Parse()
	totalIterations := *totalIterationsFlag

	for iteration := 0; iteration < totalIterations; iteration++ {
		start := time.Now()
		for i := 0; i < 2; i++ {
			var wg sync.WaitGroup
			filename := "test.zip"
			wg.Add(1)
			go downloadFile(urls, filename, &wg)
			wg.Wait()

		}
		elapsed := float32(time.Since(start).Seconds())
		fmt.Print("\033[H\033[2J")

		fmt.Printf("目前已消耗 %d GB流量！\n当前平均下载速度：%.2f Mbps\n预计还需要：%.0f 分钟\n", iteration+1, 1024/elapsed, float32(float32(totalIterations-iteration)*elapsed/60))
	}

	fmt.Printf("本次流量杀手已运行完成，总计消耗您服务器 %d GB流量！期待下次您的使用！\n", totalIterations)
}
