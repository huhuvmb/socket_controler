package main

import (
	"log"
	"net"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("消息已发送，正在关闭..."))
		cmd := exec.Command("shutdown", "/s", "/t", "1")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("消息已发送，正在重启..."))
		cmd := exec.Command("shutdown", "/r", "/t", "1")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			log.Fatal(err)
		}
		var ip string
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
				}
			}
		}
		w.Write([]byte("<html><body style=\"background-color: #f2f2f2;\"><h1 style=\"text-align: center; color: #333;\">欢迎使用远程电脑管理工具！</h1><br><br><div style=\"text-align: center;\"><a href=\"/shutdown\" onclick=\"return confirm('您确定要关闭计算机吗？')\"><button style=\"background-color: #ff4d4d; color: #fff; border: none; padding: 10px 20px; border-radius: 5px;\">关闭计算机</button></a><br><br><a href=\"/restart\" onclick=\"return confirm('您确定要重启计算机吗？')\"><button style=\"background-color: #4da6ff; color: #fff; border: none; padding: 10px 20px; border-radius: 5px;\">重启计算机</button></a></div><br><br><div style=\"text-align: center;\">被控IP地址：" + ip + "</div></body></html>"))

	})

	log.Fatal(http.ListenAndServe(":5213", nil))
}
