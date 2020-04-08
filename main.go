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
	writer.Header().Set("Access-Control-Allow-Origin", "*") ;
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type");
	writer.Header().Set("content-type", "application/json");
	server := st_server.NewStreamnetService()
	response := server.GetDagMap(nil,GatewayUrl);
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		fmt.Println(err)
	}
}

func GetTotalOrder(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*") ;
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type");
	writer.Header().Set("content-type", "application/json");
	server := st_server.NewStreamnetService()
	response := server.GetTotalOrder(nil,GatewayUrl);
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		fmt.Println(err)
	}
}


func GetBlockContent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*") ;
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type");
	writer.Header().Set("content-type", "application/json");
	server := st_server.NewStreamnetService()
	var detailRequest *st_server.NodeDetailRequest;
	if err := json.NewDecoder(request.Body).Decode(&detailRequest); err != nil {
		fmt.Println(err)
		request.Body.Close()
	}
	response := server.GetBlockContent(detailRequest,GatewayUrl);
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
	http.HandleFunc("/getTotalOrder", GetTotalOrder);
	http.HandleFunc("/getBlockContent", GetBlockContent);
	fmt.Println("启动端口" + port);
	err := http.ListenAndServe(":"+port, nil);
	if err != nil {
		fmt.Println(err);
	}
}
