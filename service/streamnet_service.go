package service

import (
	"fmt"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"github.com/triasteam/streamnet-api/common"
)

func NewStreamnetService() *server {
	return &server{}
}

type server struct {
}
type Response struct {
	Dag string `json:"dag"`
}

type TotalResponse struct {
	TotalOrder string `json:"totalOrder"`
}

func (serv *server) GetDagMap(ctx context.Context,gateway string) *common.CommonResponse {
	data := "{\"command\":\"getDAG\",\"type\":\"JSON\"}";
	r, err := doPost(gateway, []byte(data));
	if err != nil {
		return createErrorResponse(err);
	}
	var result Response;
	err = json.Unmarshal(r, &result);
	var resultString = result.Dag;
	var dagMap = map[string][]string{}
	err = json.Unmarshal([]byte(resultString),&dagMap);
	var length  = 0;
	for _,v := range dagMap {
		for range v{
			length ++
		}
	}
	if length > 0 {
		returnData := make([]map[string]string,length);
		var index = 0;
		for k,v := range dagMap {
			for i := range v{
				var unit = make(map[string]string)
				unit["source"] = k[0:6];
				unit["target"] = v[i][0:6];
				returnData[index] = unit;
				index ++;
			}
		}
		return createSuccessResponse(returnData);
	} else {
		return createErrorResponse(errors.New("No data now"))
	}

}

func (serv *server) GetTotalOrder(ctx context.Context,gateway string) *common.CommonResponse {
	data := "{\"command\":\"getTotalOrder\"}";
	r, err := doPost(gateway, []byte(data));
	if err != nil {
		return createErrorResponse(err);
	}
	var result TotalResponse;
	err = json.Unmarshal(r,&result);
	var resp string;
	resp = result.TotalOrder;
	fmt.Println(resp);
	if err != nil {
		return createSuccessResponse(result);
	} else {
		return createErrorResponse(err);
	}
	// var resultR Response;
	// err = json.Unmarshal(r, &resultR);
	// var result Response;
	// err = json.Unmarshal(r, &result);
	// var resultString = result.Dag;
	// var dagMap = map[string][]string{}
	// err = json.Unmarshal([]byte(resultString),&dagMap);
	// var length  = 0;
	// for _,v := range dagMap {
	// 	for range v{
	// 		length ++
	// 	}
	// }
	// if length > 0 {
	// 	returnData := make([]map[string]string,length);
	// 	var index = 0;
	// 	for k,v := range dagMap {
	// 		for i := range v{
	// 			var unit = make(map[string]string)
	// 			unit["source"] = k[0:6];
	// 			unit["target"] = v[i][0:6];
	// 			returnData[index] = unit;
	// 			index ++;
	// 		}
	// 	}
	// 	return createSuccessResponse(returnData);
	// } else {
	// 	return createErrorResponse(errors.New("No data now"))
	// }

}

func createErrorResponse(err error) (*common.CommonResponse) {
	resp := &common.CommonResponse{
		Code:    0,
		Data:    err.Error(),
		Message: "fail",
	}
	return resp;
}

func createSuccessResponse(data interface{}) (*common.CommonResponse) {
	resp := &common.CommonResponse{
		Code:    1,
		Data:    data,
		Message: "success",
	}
	return resp;
}

func doPost(uri string, d []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(d))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-IOTA-API-Version", "1")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return r, nil
}