package common

import (
	"io"
	"log"
	"net/http"
	"testing"
)

func TestDH_Corp_OrderPager(t *testing.T) {
	// 设置请求处理函数
	http.HandleFunc("/family", proxyHandler)

	// 启动 HTTP 服务器
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 代理处理函数
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// 要转发的目标地址
	targetURL := "https://sinotoookit.top:6443/sub?target=clash&list=true&url=https%3A%2F%2Fcloud.tpc.re%2Fapi%2Fv1%2Fclient%2Fsubscribe%3Ftoken%3D97089aa61316264ea2c24f27c6b2f068"

	// 创建转发请求
	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// 复制原始请求的头部信息
	//req.Header = r.Header

	// 执行转发请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 复制响应的头部信息
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// 设置响应状态码
	w.WriteHeader(resp.StatusCode)

	// 复制响应体
	io.Copy(w, resp.Body)
}
