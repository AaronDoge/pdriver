package main

import (
	"fmt"
	"net/http"
)

func main() {

	str := "app_id=www.tenxun.com&platform=ios&src_ip=192.168.152.12&TTL=3&dst_ip=192.168.153.218"

	byteArr := []byte(str)
	fmt.Println(len(byteArr))

	http.Post()
}

// base64
//func main() {
//	var str = "nihaoshijie"
//	strByte := []byte(str)
//	fmt.Println("编码前字节数组长度：", len(strByte))
//
//	encoded := base64.StdEncoding.EncodeToString(strByte)
//
//	encodedByte := []byte(encoded)
//	fmt.Println("编码后字节数组长度：", len(encodedByte))
//}
//编码前字节数组长度： 11
//编码后字节数组长度： 16
