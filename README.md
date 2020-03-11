[中文介绍](https://github.com/triasteam/streamnet-api/blob/master/README_cn.md)
### Streamnet-Api

&emsp;&emsp;Streamnet-Api supplies service for Streamnet-Web.Deployed on the Internet server.


##### Startup

    go build main.go
    ./main --port 7001 --gateway http://gatewayAddr:port

##### Parameter Description

    port: application listens(default:7001)
    gateway: the gateway address(proxy and balance server)

##### API Description:

Address: http://server_address/getDagMap
RequestType: GET
RequestData: none
ResponseData:

    {
      "code": 1,
      "message": "success",
      "data": [
        {
          "source": "ESRBMD",
          "target": "QPWWIH"
        },
        {
          "source": "PZ9SPU",
          "target": "9EWIUJ"
        },
        {
          "source": "PZ9SPU",
          "target": "YO9EVH"
        },
        {
          "source": "QPWWIH",
          "target": "PZ9SPU"
        }
      ]
    }


|Param|Desc|
|----|----|
|code|1:success，0:fail|
|message|success message or error infomation|
|data|data by json string|
