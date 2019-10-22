package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	st_server "github.com/triasteam/streamnet-api/service"
)

var port string
var GatewayUrl string

func GetDagMap(writer http.ResponseWriter, request *http.Request) {
	server := st_server.NewStreamnetService()
	response := server.GetDagMap(nil,GatewayUrl);
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		fmt.Println(err)
	}
}

func init() {
	flag.StringVar(&port, "port", "7001", "server port")
	flag.StringVar(&GatewayUrl, "gateway", "http://172.31.18.190:9000", "gateway url")
}

func main() {
	flag.Parse()
	if GatewayUrl == ""{
		return;
	}
	http.HandleFunc("/getDagMap", GetDagMap);
	fmt.Println("启动端口" + port);
	err := http.ListenAndServe(":"+port, nil);
	if err != nil {
		fmt.Println(err);
	}
}
