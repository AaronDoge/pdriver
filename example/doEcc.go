package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"os"
	"pdriver/testdata"
)

func main() {
	msg, _ := json.Marshal(doTest.TestBody)
	pubKeyCurve := elliptic.P256()
	stdPriKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println("生成私钥错误：", err.Error())
		os.Exit(1)
	}

	priKey := ecies.ImportECDSA(stdPriKey)

	// 公私钥编码
	//priKeyByte, err := json.Marshal(stdPriKey)
	priKeyByte, err := x509.MarshalECPrivateKey(stdPriKey)
	priKeyStr := base64.StdEncoding.EncodeToString(priKeyByte)
	fmt.Println("生成私钥：", priKeyStr)

	//pubKeyByte, err := json.Marshal(stdPriKey.PublicKey)
	pubKeyByte, err := x509.MarshalPKIXPublicKey(stdPriKey.Public())
	pubKeyStr := base64.StdEncoding.EncodeToString(pubKeyByte)
	fmt.Println("生成公钥：", pubKeyStr)

	//解码密钥对
	dePriKeyByte, err := base64.StdEncoding.DecodeString(priKeyStr)
	dePriKey, err := x509.ParseECPrivateKey(dePriKeyByte)
	priKey = ecies.ImportECDSA(dePriKey)

	dePubKeyByte, err := base64.StdEncoding.DecodeString(pubKeyStr)
	dePubKeyTmp, err := x509.ParsePKIXPublicKey(dePubKeyByte)

	dePubKey := dePubKeyTmp.(*ecdsa.PublicKey)
	fmt.Println(dePubKey)
	//

	cipherTxt, err := ecies.Encrypt(rand.Reader, &priKey.PublicKey, msg, nil, nil)

	fmt.Println("加密后密文：", base64.StdEncoding.EncodeToString(cipherTxt))

	plainTxt, err := priKey.Decrypt(cipherTxt, nil, nil)
	fmt.Println("解密后数据：", string(plainTxt))
}

