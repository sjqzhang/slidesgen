package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed index.html right.html slides.html img
var staticFS embed.FS

func main() {

	//增加命令行参数 --port，默认为8000
	//增加命令行参数 --dir 默认为当前目录

	port := flag.String("port", "8000", "服务启动端口")
	dir := flag.String("dir", ".", "指定目录")
	flag.Parse()
	// 处理特殊内置静态文件

	var fileNames = []string{
		"blackboard.png",
		"boardmarker-green.png",
		"boardmarker-red.png",
		"chalk-green.png",
		"chalk-red.png",
		"sponge.png",
		"boardmarker-black.png",
		"boardmarker-orange.png",
		"boardmarker-yellow.png",
		"chalk-orange.png",
		"chalk-white.png",
		"whiteboard.png",
		"boardmarker-blue.png",
		"boardmarker-purple.png",
		"chalk-blue.png",
		"chalk-purple.png",
		"chalk-yellow.png",
		"slidesgen-help.png",
		"2EYFMXv.woff2",
		"2yxzQba.eot",
	}

	for _, fn := range fileNames {
		fileName := fn
		http.HandleFunc(fmt.Sprintf("/img/%s", fileName), func(w http.ResponseWriter, r *http.Request) {
			c, _ := staticFS.ReadFile(fmt.Sprintf("img/%s", fileName))
			w.Write(c)
		})
	}

	// 当文件不存在时，读取指定目录下的的相应文件

	http.HandleFunc("/right.html", func(w http.ResponseWriter, r *http.Request) {
		c, _ := staticFS.ReadFile("right.html")
		w.Write(c)
	})
	http.HandleFunc("/slides.html", func(w http.ResponseWriter, r *http.Request) {
		c, _ := staticFS.ReadFile("slides.html")
		w.Write(c)
	})

	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		c, _ := staticFS.ReadFile("index.html")
		w.Write(c)
	})

	if _, err := os.Stat("index.html"); os.IsNotExist(err) {
		c, _ := staticFS.ReadFile("index.html")
		os.WriteFile("index.html", c, 0666)
	}

	http.HandleFunc("/Api/upload", uploadHandler)

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	fmt.Printf("Starting server on port %s...\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//用当前年月日时分秒生成文件名
	fileName := time.Now().Format("20060102150405")
	//以日期为目录名
	dirName := fileName[0:8]
	fmt.Println(dirName)
	//判断目录是否存在，不存在则创建
	if _, err := os.Stat("images/" + dirName); os.IsNotExist(err) {
		os.MkdirAll("images/"+dirName, 0755)
	}

	// 限制请求方法
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 解析请求体
	err := r.ParseMultipartForm(10 << 20) // 设置最大内存限制为 10 MB
	if err != nil {
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}
	// 获取文件
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error getting file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// 读取文件内容
	buf := make([]byte, header.Size)
	_, err = io.ReadFull(file, buf)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// 创建目标文件
	dst, err := os.Create(fmt.Sprintf("images/%s/%s", dirName, fileName))
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	// 写入文件内容
	_, err = dst.Write(buf)
	if err != nil {
		http.Error(w, "Error writing file", http.StatusInternalServerError)
		return
	}
	// 返回成功响应
	w.Write([]byte(fmt.Sprintf(`{"state": true, "data": "/images/%s/%s"}`, dirName, fileName)))
}
