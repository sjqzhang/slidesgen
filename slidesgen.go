package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:embed index.html right.html slides.html
var staticFS embed.FS

func main() {
	// 注册静态文件处理器
	http.Handle("/", http.FileServer(http.FS(staticFS)))
	// get port from command line
	port := ":8000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
