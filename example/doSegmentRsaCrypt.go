package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"pdriver/encrypt"
	"pdriver/testdata"
	"time"
)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDi0zTKSGQMVuhlBfahzHyspPmL
aPZvEdusUq8o0vOPxnDQwSET52W2n9Fv+L/PTuzfKhaFpqvZf03TLu2IDheYOips
Bu7KB/lGJgoVE8hgn+G5V5JotUZ4/u30GGTV3MYMTyJHcgS3KDrx/mKCMd+1Gr9u
g63WmbVoCKHjHgIccQIDAQAB
-----END PUBLIC KEY-----
`)

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

var test_data = &doTest.ReqBody{
	Token: 	"",
	Action: 	"ping",
	Timestamp: 	1545733726,
	IPInfo: doTest.IP_Info{
		City: "shanghai_yangpu",
		Country: "china",
		IP: 	"192.168.152.12",
		Location: "string",
		Org: 	"string",
		Region: "shanghai",
	},
	PingData: doTest.Ping_Data{
		TTL: "3",
		Delay: 1,
		DstIP: "192.168.153.218",
		Loss: 0,
		SrcIP: "192.168.152.12",
	},
}

func main() {
	tagCipher := "816300430f51b96bd442a3b8574e66026a03d338e22b428860a1c0a2383d8156e7e593ef7636095e9fd123afb9291169a472fb67785e48145ff75dd607ee6e745fce90bae007cec319bec60e1ba11b84cfc17840e407f5b02577803140193ebdb313928c8b7a7a1bbca9c64fcb99e225f956a9f575506987a2261614017f2870"
	//tagByte, err := hex.DecodeString(tagCipher)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	tagtmp, err := encrypt.SegmentDecrypt(tagCipher, privateKey2)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("解密后：", string(tagtmp))

	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)

	rand.Seed(time.Now().Unix() - int64(1))
	randNum2 := rand.Intn(100)

	fmt.Printf("rand is %v\n", randNum)
	fmt.Printf("rand is %v\n", randNum2)

}

func main1() {

	//startEnTime := time.Now()
	//startEnTimeN := time.Now().Nanosecond()
	//fmt.Println("打印Unix时间：",startEnTimeN)
	//fmt.Println("开始加密时间：", startEnTime)
	// 先加密
	dataByte, _ := json.Marshal(test_data)
	//var dataSeg  []byte
	//for i := 0; i < len(dataByte); i += 117 {
	//	var seg []byte
	//	if i+117 < len(dataByte) {
	//		seg = dataByte[i:i+117]
	//	} else {
	//		seg = dataByte[i:len(dataByte)]
	//	}
	//	dataSeg = append(dataSeg, seg...)
	//}
	starto := time.Now()
	cipherData, err := encrypt.SegmentEncrypt(dataByte, publicKey)
	if err != nil {
		fmt.Println("加密失败：", err.Error())
	}


	//fmt.Println("加密后密文：", cipherData)
	//fmt.Println("加密后密文总长度：", len(cipherData))
	//endEnTime := time.Now()
	//fmt.Println("加密结束时间：", endEnTime)
	//fmt.Println("加密完成时间：", endEnTime.Sub(startEnTime))

	//startDeTime := time.Now()
	//startDeTimeN := time.Now().Nanosecond()
	//fmt.Println("开始解密时间：", startDeTime)
	// 使用加密密文解密
	// 密文长度检查，不等于0，错误
	//fmt.Println("密文长度检查：", len(cipherData) % 172)

	plainDataByte, err := encrypt.SegmentDecrypt(cipherData, privateKey2)
	if err != nil {
		fmt.Println("解密失败", err.Error())
	}
	//plainData := base64.StdEncoding.EncodeToString(plainDateByte)
	plainData := string(plainDataByte)

	endtimeo := time.Now()
	fmt.Println("解密后数据：", plainData)

	//endDeTime := time.Now()
	//endDeTimeN := time.Now().Nanosecond()
	//fmt.Println("解密完成时间：", endDeTime)
	//fmt.Println("解密用时(ns)：", endDeTimeN - startDeTimeN)
	//fmt.Println("解密用时(ms)：", endDeTime.Sub(startDeTime))
	fmt.Println("加解密总用时：", endtimeo.Sub(starto))
}



func main2() {

	startEnTime := time.Now()
	startEnTimeN := time.Now().Nanosecond()
	fmt.Println("打印Unix时间：",startEnTimeN)
	fmt.Println("开始加密时间：", startEnTime)
	// 先加密
	dataByte, _ := json.Marshal(test_data)
	var dataSeg  [][]byte
	for i := 0; i < len(dataByte); i += 117 {
		var seg []byte
		if i+117 < len(dataByte) {
			seg = dataByte[i:i+117]
		} else {
			seg = dataByte[i:len(dataByte)]
		}
		dataSeg = append(dataSeg, seg)
	}
	var cipherData string

	for i := 0; i < len(dataSeg); i++ {
		cipherDataByte, err := encrypt.RsaEncrypt(dataSeg[i], publicKey)
		if err != nil {
			fmt.Println("加密失败：", err.Error())
		}
		cipherDataTmp := hex.EncodeToString(cipherDataByte)

		cipherData += cipherDataTmp
	}

	fmt.Println("加密后密文：", cipherData)
	fmt.Println("加密后密文总长度：", len(cipherData))
	endEnTime := time.Now()
	fmt.Println("加密结束时间：", endEnTime)
	fmt.Println("加密完成时间：", endEnTime.Sub(startEnTime))


	startDeTime := time.Now()
	startDeTimeN := time.Now().Nanosecond()
	fmt.Println("开始解密时间：", startDeTime)
	// 使用加密密文解密
	// 密文长度检查，不等于0，错误
	fmt.Println("密文长度检查：", len(cipherData) % 256)

	var deCipherDataArr []string
	deCipherDataTmp := []rune(cipherData)
	for i := 0; i < len(cipherData); i+=256 {
		deCipherDataArr = append(deCipherDataArr, string(deCipherDataTmp[i:i+256]))
	}

	var deCipherByte []byte
	for i := 0; i < len(deCipherDataArr); i++ {
		fmt.Println("密文数组：", i,  deCipherDataArr[i])

		deCipherByteTmp, err := hex.DecodeString(deCipherDataArr[i])
		if err != nil {
			fmt.Println("decode string to []byte error", err.Error())
		}

		deCipherDataByte, err := encrypt.RsaDecrypt(deCipherByteTmp, privateKey2)
		if err != nil {
			fmt.Println("decrypt error", err.Error())
		}

		deCipherByte = append(deCipherByte, deCipherDataByte...)
	}

	deCipherData := string(deCipherByte)

	fmt.Println("解密后数据：", deCipherData)

	endDeTime := time.Now()
	endDeTimeN := time.Now().Nanosecond()
	fmt.Println("解密完成时间：", endDeTime)
	fmt.Println("解密用时(ns)：", endDeTimeN - startDeTimeN)
	fmt.Println("解密用时(ms)：", endDeTime.Sub(startDeTime))
}
