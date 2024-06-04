package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
)

//go:embed index.html right.html slides.html
var staticFS embed.FS

func main() {

	//增加命令行参数 --port，默认为8000
	//增加命令行参数 --dir 默认为当前目录

	port := flag.String("port", "8000", "服务启动端口")
	dir := flag.String("dir", ".", "指定目录")
	flag.Parse()
	// 注册静态文件处理器

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

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	fmt.Printf("Starting server on port %s...\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}
