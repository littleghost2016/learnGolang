package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>This is my third page</h1>")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
	// 网站上云注意安全组放行对应端口
}
