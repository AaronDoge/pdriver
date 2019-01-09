package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"time"
)

type ReqBody struct {
	Token 	string	`json:"token"`
	Action 	string	`json:"action"`
	Timestamp 	int	`json:"timestamp"`
	IPInfo 	IP_Info	`json:"ip_info"`
	PingData Ping_Data	`json:"ping_data"`
}
type IP_Info struct {
	City 		string	`json:"city"`
	Country 	string	`json:"country"`
	IP 			string		`json:"ip"`
	Location 	string	`json:"location"`
	Org 		string		`json:"org"`
	Region 		string		`json:"region"`
}
type Ping_Data struct {
	TTL   string `json:"TTL"`
	Delay int `json:"delay"`
	DstIP string `json:"dst_ip"`
	Loss  int    `json:"loss"`
	SrcIP string    `json:"src_ip"`
}

// testdata



var rBody = &ReqBody{
	Token: 	"",
	Action: 	"ping",
	Timestamp: 	1545733726,
	IPInfo: IP_Info{
		City: "shanghai_yangpu",
		Country: "china",
		IP: 	"192.168.152.12",
		Location: "string",
		Org: 	"string",
		Region: "shanghai",
	},
	PingData: Ping_Data{
		TTL: "3",
		Delay: 1,
		DstIP: "192.168.153.218",
		Loss: 0,
		SrcIP: "192.168.152.12",
	},
}

var pingData = &Ping_Data{
	TTL: "3",
	Delay: 1,
	DstIP: "192.168.153.218",
	Loss: 0,
	SrcIP: "192.168.152.12",
}

var ipInfo = &IP_Info{
	City: "shanghai_yangpu_longchanglu",
	Country: "china",
	IP: 	"192.168.152.12",
	Location: "string",
	Org: 	"string",
	Region: "shanghai",
}


var (
	DefaultCurve = ethcrypto.S256()
)

func main() {
	start := time.Now()
	pubCurve := elliptic.P256()
	ecdsPriKey, err := ecdsa.GenerateKey(pubCurve, rand.Reader)
	if err != nil {
		fmt.Println("生成密钥错误", err.Error())
		return
	}

	//privateKey, err := ecies.GenerateKey(rand.Reader, DefaultCurve, nil)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//privateKey2, err := ecies.GenerateKey(rand.Reader, DefaultCurve, nil)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	eciesPriKey := ecies.ImportECDSA(ecdsPriKey)
	//eciesPubKey := eciesPriKey.PublicKey

	priKeyStr, err := x509.MarshalECPrivateKey(ecdsPriKey)
	if err != nil {
		fmt.Println("私钥编码错误", err.Error())
		return
	}
	fmt.Println("私钥：", base64.StdEncoding.EncodeToString(priKeyStr))

	pubKeyStr, err := x509.MarshalPKIXPublicKey(eciesPriKey.PublicKey.ExportECDSA())
	if err != nil {
		fmt.Println("公钥编码错误", err.Error())
		return
	}
	fmt.Println("公钥x509编码：", base64.StdEncoding.EncodeToString(pubKeyStr))



	//message := []byte("hello, world")
	//message, _ := json.Marshal(rBody)
	message := []byte("</2018>2018努力让自己不一样，却始终没能跳出圈子；对未来充满忧虑；2019 make a change！如果你看到，与你共勉！<2019>")
	cipherTxt, err := ecies.Encrypt(rand.Reader, &eciesPriKey.PublicKey, message, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println("密文：", string(cipherTxt))
	cipherTxtStr := base64.StdEncoding.EncodeToString(cipherTxt)
	fmt.Println("base64编码密文：", cipherTxtStr, "\n长度：", len(cipherTxtStr))


	plainTxt, err := eciesPriKey.Decrypt(cipherTxt, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	end := time.Now()
	fmt.Println("加解密总用时：", end.Sub(start))
	fmt.Println("解密后数据：", string(plainTxt))

	if !bytes.Equal(plainTxt, message) {
		fmt.Println("解密失败！解密信息和源数据不符！")
		return
	}

	_, err = eciesPriKey.Decrypt(cipherTxt, nil, nil)
	if err != nil {
		fmt.Println("ecies: encryption should not have succeeded")
		return
	}
}

