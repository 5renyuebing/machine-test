package main

import (
	"encoding/base64"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const s = "00000008d78d46a"

func params_b64(path string, params map[string]string) string {
	params_list := []string{}
	for _, v := range params {
		params_list = append(params_list, v)
	}
	sort.Sort(sort.StringSlice(params_list))
	params_string := strings.Replace(strings.Trim(fmt.Sprint(params_list), "[]"), " ", "", -1)
	params_b64 := base64.StdEncoding.EncodeToString([]byte(params_string))
	t := int(time.Now().Unix() * 1000) - 1515125653845
	res := params_b64 + "@#" + path + "@#" + strconv.Itoa(t) + "@#1"
	return res
}

func data_encrypt(data string) string {
	data_list := []byte{}
	for _, v := range data{
		data_list = append(data_list, byte(v))
	}
	for k, v := range data_list {
		data_list[k] = byte(int32(v) ^ int32(s[k % len(s)]))
	}
	res := base64.StdEncoding.EncodeToString(data_list)
	return res
}

func main()  {
	m := map[string]string{}
	m["appid"] = "17"
	path := "/andapp/comment"
	h := map[string]string{}
	h["Accept"] = "application/json, text/plain, */*"
	h["Origin"] = "https: // www.qimai.cn"
	h["Referer"] = "https://www.qimai.cn/andapp/comment/appid/1935455"
	h["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"
	res := params_b64(path, m)
	analysis := data_encrypt(res)
	url := "https://api.qimai.cn" + path + "?appid=17&analysis=" + analysis
}
