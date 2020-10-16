package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	SERVER1 = "https://www.google.com"
	SERVER2 = "https://www.reddit.com"
	SERVER3 = "https://www.yahoo.com"
	PORT = "1338"
)
var serverURLs = [] string{SERVER1,SERVER2,SERVER3}
var serverCount = 0

func main()  {
	http.HandleFunc("/",loadBalancer)
	http.ListenAndServe(":"+PORT,nil)
}

func loadBalancer(res http.ResponseWriter, req *http.Request)  {
	url := getProxyURL()
	logRequestPayload(url,req)
	serveReverseProxy(url,res,req)
}

func getProxyURL()  string {
	server := serverURLs[serverCount]
	serverCount++
	if serverCount >= len(serverURLs){
		serverCount = 0
	}
	return server
}

func serveReverseProxy(target string ,res http.ResponseWriter, req *http.Request)  {
	url,_ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res,req)
}

func logRequestPayload(url string,req *http.Request)  {
	fmt.Println("url :",url)
}