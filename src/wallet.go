package src

import (
	_"bian/config"
	"bian/net"
	"fmt"
	"strconv"
	"time"
)

//获取系统状态
func GetSystemState()  {
	path := "/sapi/v1/system/status"
	res, err := net.GetRequest(path,nil)

	if err != nil{
		fmt.Println("请求失败：",err)
	}

	fmt.Println("res = ",res)
}

//获取所有币种信息
func GetAllCoinInfo()  {
	path := "/sapi/v1/capital/config/getall"
	params := make(map[string]string)
	params["recvWindow"] = "5"
	params["timestamp"] = strconv.Itoa(int(time.Now().Unix()))

	//res := net.HttpGetRequest(path,params)
	//fmt.Println(res)
	net.GetRequest(path,params)
}