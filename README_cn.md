### Streamnet-Api
&emsp;&emsp;Streamnet-Api 为Streamnet-Web提供服务。部署在外网服务器上。

启动方法 

    go build main.go
    ./main --port 7001 --gateway http://gatewayAddr:port

参数说明
port 启动端口（默认7001）
gateway 网关地址

#### 接口说明:
地址: http://server_address/getDagMap

请求方式: GET

请求数据: 无

响应数据: 

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

说明

|字段|说明|
|----|----|
|code|状态码1成功，0失败|
|message|与code对应的消息|
|data|具体数据|
