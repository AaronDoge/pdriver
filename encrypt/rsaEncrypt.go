package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
)

/*
加密长度限制
解决方案1，分块加密

 */
func RsaEncrypt(plainData, publicKey []byte) ([]byte, error){
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainData)
}

func RsaDecrypt(cipherData, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, pri, cipherData)
}

// 分段加密
func SegmentEncrypt(plainData []byte, publicKey []byte) (string, error) {

	var dataSeg  [][]byte
	for i := 0; i < len(plainData); i += 117 {
		var seg []byte
		if i+117 < len(plainData) {
			seg = plainData[i:i+117]
		} else {
			seg = plainData[i:len(plainData)]
		}
		dataSeg = append(dataSeg, seg)
	}

	var cipherData string
	for i := 0; i < len(dataSeg); i++ {
		cipherDataByte, err := RsaEncrypt(dataSeg[i], publicKey)
		if err != nil {
			fmt.Println("加密失败：", err.Error())
			return "", err
		}
		cipherDataTmp := hex.EncodeToString(cipherDataByte)

		cipherData += cipherDataTmp
	}

	return cipherData, nil
}

// 分段加密
func SegmentEncrypt2(plainData []byte, publicKey []byte) (string, error) {

	var dataSeg  [][]byte
	for i := 0; i < len(plainData); i += 117 {
		var seg []byte
		if i+117 < len(plainData) {
			seg = plainData[i:i+117]
		} else {
			seg = plainData[i:len(plainData)]
		}
		dataSeg = append(dataSeg, seg)
	}

	var cipherByteArr []byte
	for i := 0; i < len(dataSeg); i++ {
		cipherDataByte, err := RsaEncrypt(dataSeg[i], publicKey)

		if err != nil {
			fmt.Println("加密失败：", err.Error())
			return "", err
		}
		fmt.Printf("第%d段加密字节数组长度%d, 字节数组：%s\n", i, len(cipherDataByte), cipherDataByte)
		cipherByteArr = append(cipherByteArr, cipherDataByte...)
	}

	//cipherDataTmp := hex.EncodeToString(cipherDataByte)
	cipherData := base64.StdEncoding.EncodeToString(cipherByteArr)


	return cipherData, nil
}


// 分段解密
func SegmentDecrypt(cipherData string, privateKey []byte) ([]byte, error) {
	if len(privateKey) == 0 {
		return nil, errors.New("私钥错误")
	}
	// 密文长度检查
	if len(cipherData) % 256 != 0 {
		fmt.Println("密文检查错误，错误的密文长度")
		return nil, errors.New("密文检查错误，错误的密文长度")
	}

	// 解析出每一段密文，分别解密
	var deCipherDataArr []string
	deCipherDataTmp := []rune(cipherData)
	for i := 0; i < len(cipherData); i+=256 {
		deCipherDataArr = append(deCipherDataArr, string(deCipherDataTmp[i:i+256]))
	}

	var deCipherByte []byte
	for i := 0; i < len(deCipherDataArr); i++ {
		deCipherByteTmp, err := hex.DecodeString(deCipherDataArr[i])
		//deCipherByteTmp, err :=base64.StdEncoding.DecodeString(deCipherDataArr[i])
		if err != nil {
			fmt.Println("decode string to []byte error", err.Error())
			return nil, err
		}
		deCipherDataByte, err := RsaDecrypt(deCipherByteTmp, privateKey)
		if err != nil {
			fmt.Println("decrypt error", err.Error())
			return  nil, err
		}

		deCipherByte = append(deCipherByte, deCipherDataByte...)
	}

	return deCipherByte, nil
}
