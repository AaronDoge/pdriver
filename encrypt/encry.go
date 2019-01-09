package encrypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"goEncrypt"
	"os"
	"time"
)

func Encrypt()  {
	plaintext := []byte("床前明月光，疑是地上霜，举头望明月，学习go语言")//明文
	fmt.Println("明文为：",string(plaintext))

	//传入明文和自己定义的密钥，密钥为8字节，如果不足8字节函数内部自动补全，超过8字节函数内部截取
	cryptText := goEncrypt.DesCBC_Encrypt(plaintext, []byte("asd12345")) //得到密文
	fmt.Println("DES的CBC模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

	st := time.Now().Nanosecond()
	//传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错，8字节，如果不足8字节函数内部自动补全，超过8字节函数内部截取
	newplaintext := goEncrypt.DesCBC_Decrypt(cryptText, []byte("asd12345"))  //解密得到密文
	time.Sleep(1*time.Second)
	et := time.Now().Nanosecond()

	fmt.Println("DES开始解密时间：", st)
	fmt.Println("DES解密完成时间：", et)
	fmt.Println("DES解密时间：", et-st)

	fmt.Println("DES的CBC模式解密完：", string(newplaintext))
}

// ecc

//func EccEncrypt(rand *io.Reader, message, publicKey []byte) ([]byte, error){
<<<<<<< HEAD
//	syscall.SIGKILL
=======
>>>>>>> 97f8ad73a3270f4c5685c6018a8b867628468573
//
//}

func doEccEncrypt() {
	pubKeyCurve := elliptic.P256()
	priKey, err := ecdsa.GenerateKey(pubKeyCurve, rand.Reader)
	if err != nil {
		fmt.Println("生成私钥错误：", err.Error())
		os.Exit(1)
	}

	fmt.Println("生成私钥：", priKey)
	fmt.Println("生成公钥：", priKey.PublicKey)
}
