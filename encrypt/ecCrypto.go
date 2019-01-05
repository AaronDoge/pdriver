package encrypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

func EccEncrypt() {
	pubKeyCurve := elliptic.P256()
	priKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println("生成私钥错误：", err.Error())
		os.Exit(1)
	}

	fmt.Println("生成私钥：", priKey)
	fmt.Println("生成公钥：", priKey.PublicKey)
}
