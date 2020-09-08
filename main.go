package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	host         = ""
)
func GetServerIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetClientIp(r *http.Request) string {
	clientIP := r.RemoteAddr
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		clientIP = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		clientIP = ip
	} else {
		clientIP, _, _ = net.SplitHostPort(clientIP)
	}
	if clientIP == "::1" {
		clientIP = "127.0.0.1"
	}
	return clientIP
}

func main() {
	flag.StringVar(&host, "host", "0.0.0.0:38001", "")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		rs, _:= json.Marshal(map[string]interface{}{
			"ServerIP": GetServerIP(),
			"ClientIP": GetClientIp(r),
			"Version": "latest",
			"Host": hostname,
			"Time": time.Now().Format("2006-01-02 15:04:05"),
		})
		w.Write([]byte(rs))
	})
	fmt.Printf("The k8s demo is running at http://%s\n", host)
	fmt.Printf("Quit the server with Control-C\n\n")
	if err := http.ListenAndServe(host, nil); err != nil {
		fmt.Print(err)
	}
}
