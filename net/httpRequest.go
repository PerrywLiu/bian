package net

import (
	"bian/config"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetRequest(path string, param map[string]string) (interface{},error) {
	client := &http.Client{}
	url := config.Host + path
	query := searielParams(param)
	fmt.Println("url = ",url+query)

	//header.Set("")
	request,err := http.NewRequest("GET",url+query,nil)

	if err != nil{
		fmt.Println(path + " 请求失败：",err)
		return nil,err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	request.Header.Set("X-MBX-APIKEY",config.APIKey)

	res, err := client.Do(request)
	defer res.Body.Close()

	body,err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println(path + " 解析失败：",err)
		return nil,err
	}
	fmt.Println("res = ",string(body))
	return string(body),nil
}

func HttpGetRequest(strUrl string, mapParams map[string]string) string {
	httpClient := &http.Client{}

	var strRequestUrl string
	if nil == mapParams {
		strRequestUrl = strUrl
	} else {
		strParams := Map2UrlQuery(mapParams)
		strRequestUrl = strUrl + "?" + strParams
	}

	// 构建Request, 并且按官方要求添加Http Header
	request, err := http.NewRequest("GET", strRequestUrl, nil)
	if nil != err {
		return err.Error()
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	request.Header.Set("X-MBX-APIKEY",config.APIKey)
	// 发出请求
	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error()
	}
	defer response.Body.Close()
	// 解析响应内容
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}

func searielParams(params map[string]string) string {
	var res string = ""
	for key, value := range params{
		res += (key + "=" + value + "&")
	}
	if len(res) > 0 {
		//res = "apikey="+config.APIKey + "&secrectkey="+config.SecretKey
		res = strings.TrimSuffix(res,"&")
		res = "?" + res
	}
	return res
}


// 将map格式的请求参数转换为字符串格式的
// mapParams: map格式的参数键值对
// return: 查询字符串
func Map2UrlQuery(mapParams map[string]string) string {
	var strParams string
	for key, value := range mapParams {
		strParams += (key + "=" + value + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}
