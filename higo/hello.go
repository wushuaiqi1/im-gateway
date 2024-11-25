package higo

import "net/http"

// hello 服务测试接口

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("你好"))
	if err != nil {
		return
	}
}
