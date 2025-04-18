package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
    "time"
)

// CPU密集型任务
func cpuIntensiveTask() {
	for i := 0; i < 1000000000; i++ {
		_ = i * i // 进行一些计算
	}
    for i := 0; i < 1000000000; i++ {
		_ = i * i // 进行一些计算
	}
    for i := 0; i < 1000000000; i++ {
		_ = i * i // 进行一些计算
	}
    for i := 0; i < 1000000000; i++ {
		_ = i * i // 进行一些计算
	}
    for i := 0; i < 1000000000; i++ {
		_ = i * i // 进行一些计算
	}
}

// 模拟网络请求
func networkRequest() {
	resp, err := http.Get("http://172.17.189.156:8898/help/searchProject")
	if err != nil {
		fmt.Println("网络请求错误:", err)
		return
	}
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body) // 读取响应体
	if err != nil {
		fmt.Println("读取响应体错误:", err)
		return
	}

	fmt.Println("网络请求返回结果:", string(body)) // 打印响应内容
}

// 模拟文件操作
func fileOperations() {
	file, err := os.Create("example.txt") // 打开文件
	if err != nil {
		fmt.Println("文件打开错误:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, World!") // 写文件
	if err != nil {
		fmt.Println("文件写入错误:", err)
		return
	}

	content, err := ioutil.ReadFile("example.txt") // 读取文件
	if err != nil {
		fmt.Println("文件读取错误:", err)
		return
	}
	fmt.Println("文件内容:", string(content))
}

// 示例的 helloHandler 函数
func helloHandler() {
	fmt.Println("Hello, this is a separate handler function!")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	// 启动 CPU 密集型任务
	wg.Add(1)
	go func() {
		defer wg.Done()
		cpuIntensiveTask()
	}()

	// 启动多个网络请求
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}

	// 启动文件操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		fileOperations()
	}()

	// 调用 helloHandler
	helloHandler()

	// 等待所有 goroutine 完成
	wg.Wait()

    time.Sleep(1 * time.Second)
	fmt.Fprintln(w, "所有任务完成")


}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("服务启动，监听端口 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}
