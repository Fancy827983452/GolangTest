package algorithm

import (
	"crypto/aes"
	"crypto/cipher"
	//"fmt"
	"time"
	"math/rand"
)

//对称加密
//AEC加密和解密（CRT模式）存入数据库的信息都需要经过加密再insert，解密同样调用这一函数，即可解密
func AEC_CRT_Crypt(text []byte,key []byte) []byte{
	//指定加密、解密算法为AES，返回一个AES的Block接口对象
	block,err:=aes.NewCipher(key)
	if err!=nil{
		panic(err)
	}
	//指定计数器,长度必须等于block的块尺寸16bit
	count:=[]byte("12345678abcdefgh")
	//指定分组模式
	blockMode:=cipher.NewCTR(block,count)
	//执行加密、解密操作
	message:=make([]byte,len(text))
	blockMode.XORKeyStream(message,text)
	//返回明文或密文
	return message
}
//生成AEC密钥 （随机字符串（数字和小写字母混合，l表示指定长度））一个用户只有一个密钥
func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//
//func main(){
//	message:=[]byte("Hello!My name is X.")
//	//指定密钥
//	key:=[]byte(GetRandomString(16))
//	//加密
//	cipherText:=AEC_CRT_Crypt(message,key)
//	fmt.Println("加密后为：",string(cipherText))
//	//解密
//	plainText:=AEC_CRT_Crypt(cipherText,key)
//	fmt.Println("解密后为：",string(plainText))
//}

