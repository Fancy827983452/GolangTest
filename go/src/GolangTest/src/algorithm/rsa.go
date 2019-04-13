package algorithm

import (
	"crypto/rsa"
	"encoding/pem"
	"crypto/x509"
	"os"
	"io/ioutil"
	"crypto/sha256"
	"crypto"
	"crypto/rand"
	"util"
	"errors"
)
var privateKey, publicKey []byte

//生成私钥和公钥文件，int参数为生成私钥的长度，通常为1024或2048
func GenRsaKey(bits int) error {
	//生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	util.CheckErr(err)
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	util.CheckErr(err)
	err = pem.Encode(file, block)
	util.CheckErr(err)

	//生成公钥文件
	publicKey := &privateKey.PublicKey
	defPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	util.CheckErr(err)
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: defPkix,
	}
	file, err = os.Create("public.pem")
	util.CheckErr(err)
	err = pem.Encode(file, block)
	util.CheckErr(err)
	return nil
}

func GetPublicKey(file string) ([]byte,error){//读取公钥
	publicKey, err := ioutil.ReadFile(file)
	return  publicKey,err
}

func GetPrivateKey(file string) ([]byte,error){//读取私钥
	privateKey, err := ioutil.ReadFile(file)
	return  privateKey,err
}

//私钥签名
func RsaSign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	util.CheckErr(err)
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
}

//公钥验证
func RsaSignVer(data []byte, signature []byte) error {
	hashed := sha256.Sum256(data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	util.CheckErr(err)
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// 公钥加密
func RsaEncrypt(data []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	util.CheckErr(err)
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// 私钥解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	util.CheckErr(err)
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//使用示例
//func test() {
//	var theMsg = "the message you want to encode 你好 世界"
//	fmt.Println("Source:", theMsg)
//  err:=algorithm.GenRsaKey(256) //生成256位的公钥和私钥对，保存在本地
//  util.CheckErr(err)
//
//	//获取公钥和私钥
//	publicKey, err :=GetPublicKey("public.pem")
//	CheckErr(err)
//	privateKey, err := GetPublicKey("private.pem")
//	CheckErr(err)
//	fmt.Printf("%s\n", publicKey)
//	fmt.Printf("%s\n", privateKey)
//
//	//私钥签名
//	sig, _ := RsaSign([]byte(theMsg))
//	fmt.Println(string(sig))
//	//公钥验证
//	fmt.Println(RsaSignVer([]byte(theMsg), sig))
//
//	//公钥加密
//	enc, _ := RsaEncrypt([]byte(theMsg))
//	fmt.Println("Encrypted:", string(enc))
//	//私钥解密
//	decstr, _ := RsaDecrypt(enc)
//	fmt.Println("Decrypted:", string(decstr))
//}
