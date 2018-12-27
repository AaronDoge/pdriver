package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"pdriver/encrypt"
)

type reqBody struct {
	Token 	string	`json:"token"`
	Action 	string	`json:"action"`
	Timestamp 	int	`json:"time"`
	IPInfo 	IPInfo	`json:"ip_info"`
	PingData PingData	`json:"ping_data"`
}
type IPInfo struct {
	City 	string	`json:"city"`
	Country 	string	`json:"country"`
	IP 		string		`json:"ip"`
	Location 	string	`json:"location"`
	Org 	string		`json:"org"`
	Region 	string		`json:"region"`
}
type PingData struct {
	TTL   string `json:"TTL"`
	Delay int `json:"delay"`
	DstIP string `json:"dst_ip"`
	Loss  int    `json:"loss"`
	SrcIP string    `json:"src_ip"`
}

// testdata
var testdata = &reqBody{
	Token: 	"",
	Action: 	"ping",
	Timestamp: 	1537342710,
	IPInfo: IPInfo{
		City: "shanghai",
		Country: "china",
		IP: 	"192.168.152.12",
		Location: "string",
		Org: 	"string",
		Region: "shanghai",
	},
	PingData: PingData{
		TTL: "3",
		Delay: 1,
		DstIP: "192.168.153.218",
		Loss: 0,
		SrcIP: "192.168.152.12",
	},
}

var pingdata = &PingData{
	TTL: "3",
	Delay: 1,
	DstIP: "192.168.153.218",
	Loss: 0,
	SrcIP: "192.168.152.12",
}

var ipinfo = &IPInfo{
	City: "shanghai",
	Country: "china",
	IP: 	"192.168.152.12",
	Location: "string",
	Org: 	"string",
	Region: "shanghai",
}

func main() {
	// 公钥和私钥可以从文件中读取
//	var privateKey = []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
//7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
//Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
//AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
//ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
//XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
///jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
//IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
//4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
//DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
//9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
//DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
//AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
//-----END RSA PRIVATE KEY-----
//`)

	var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDi0zTKSGQMVuhlBfahzHyspPmL
aPZvEdusUq8o0vOPxnDQwSET52W2n9Fv+L/PTuzfKhaFpqvZf03TLu2IDheYOips
Bu7KB/lGJgoVE8hgn+G5V5JotUZ4/u30GGTV3MYMTyJHcgS3KDrx/mKCMd+1Gr9u
g63WmbVoCKHjHgIccQIDAQAB
-----END PUBLIC KEY-----
`)
	var pubtmp = []byte(`
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDi0zTKSGQMVuhlBfahzHyspPmL
aPZvEdusUq8o0vOPxnDQwSET52W2n9Fv+L/PTuzfKhaFpqvZf03TLu2IDheYOips
Bu7KB/lGJgoVE8hgn+G5V5JotUZ4/u30GGTV3MYMTyJHcgS3KDrx/mKCMd+1Gr9u
g63WmbVoCKHjHgIccQIDAQAB
`)
	fmt.Println("公钥长度：", len(pubtmp))
	//data := []byte("this is a test data")
	// mp := map[string]string{"token":"", "action":"ping", "timestamp": ""}

	data, err := json.Marshal(ipinfo)
	//var dataBlock [][]byte
	//fmt.Println("字节数组长度：", len(data))
	//
	//for i := 0; i < len(data) ; i += 100 {
	//	dataBlock = append(dataBlock, data[i: i+100])
	//}
	if err != nil {
		fmt.Println("json marshal error. ", err.Error())
	}
	fmt.Println("加密前数据字节数组长度：", len(data))

	cipherData, err := encrypt.RsaEncrypt(data, publicKey)
	if err != nil {
		fmt.Println("加密失败：", err.Error())
	}

	fmt.Println("加密后数据字节数组长度：", len(cipherData))
	// 转为字符串
	cip := hex.EncodeToString(cipherData)

	fmt.Println("加密后的数据：", cip)
	fmt.Println("密文长度：", len(cip))

	var privateKey2 = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDi0zTKSGQMVuhlBfahzHyspPmLaPZvEdusUq8o0vOPxnDQwSET
52W2n9Fv+L/PTuzfKhaFpqvZf03TLu2IDheYOipsBu7KB/lGJgoVE8hgn+G5V5Jo
tUZ4/u30GGTV3MYMTyJHcgS3KDrx/mKCMd+1Gr9ug63WmbVoCKHjHgIccQIDAQAB
AoGBAJvUPhA6a4GOs1m0HxxJP93b+RStp5/mxOQ+adfCFVJRInBIPlFOR7KPWXNz
kVL7BgDLCW4Ic0eZDf5n3wf5pnrDOizkAdYQ9t0Xvd12MFGsEb3qhmt8+JPQfhgn
d/NPj9RsAskBtXmvREFWDhqlV9xtm2tlzkvbEwNKo1WjzXGBAkEA6b8mn0HS5gBk
sHv9L+NpSrTVeGCVEDVVm5/K4BR9j4GXpXM9N11G7t3nbEiwYDszuVIlwA4SxwN3
h32VRi9W6wJBAPhrXIXLkPQMRObBSCyYUlOzSuAUaHVfpOIFVTg450UP2arTT1f0
ys6ScH5c3P00amyScPeyZ7RQ+IYER53DuxMCQQCfM02+0jot5L6vZQNAhobEFv39
iup7q5eu8tpeXBZYk08RpLdg6erR7dkc6zUVlbzYz5ZehDdNzJKweVwd/UgZAkEA
iAm7HCXTFmJVpQw5avprMxzfJwDmB0i+MWv8NBKtS1uXtn2LWL5cBW2aHwjZl+uJ
UDWNmYdgVRV2U2WqllGmcwJALZPsc2mfcXjdkIuK0m+QpTQtvFzQ96sx0FTs8/qK
9Dg3Jvuvb8phc/YQTyfW0gn9Lyc4F6txUvSv8jzMC+mT2A==
-----END RSA PRIVATE KEY-----
`)
	decryData, _ := hex.DecodeString(cip)
	plainData, err := encrypt.RsaDecrypt(decryData, privateKey2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("解密后的数据：", string(plainData))

	var de IPInfo
	err = json.Unmarshal(plainData, &de)

	fmt.Println("解析到struct：", de)
	fmt.Println(de.Region, de.Org, de.Location, de.IP, de.Country, de.City)
}
