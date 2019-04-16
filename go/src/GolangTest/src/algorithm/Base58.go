package algorithm
import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	//"fmt"
	"encoding/hex"
	//"crypto/ecdsa"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
const version = byte(0x00) //定义版本号，一个字节
const addressChecksumLen = 4

// 字节数组转 Base58,加密
func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)
	for b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return result
}

// Base58转字节数组，解密
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0

	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}

	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded
}

// 字节数组反转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

//type Wallet struct {
//	//1. 私钥
//	PrivateKey ecdsa.PrivateKey
//
//	//2. 公钥
//	PublicKey  []byte
//}

//判断地址是否有效
func IsValidForAdress(adress []byte) bool {
	//将地址进行base58反编码，生成的其实是version+Pub Key hash+ checksum这25个字节
	version_public_checksumBytes := Base58Decode(adress)

	//[25-4:],就是21个字节往后的数（22,23,24,25一共4个字节）
	checkSumBytes := version_public_checksumBytes[len(version_public_checksumBytes) - addressChecksumLen:]
	//[:25-4],就是前21个字节（1～21,一共21个字节）
	version_ripemd160 := version_public_checksumBytes[:len(version_public_checksumBytes) - addressChecksumLen]
	//取version+public+checksum的字节数组的前21个字节进行两次256哈希运算，取结果值的前4个字节
	checkBytes := CheckSum(version_ripemd160)
	//将checksum比较，如果一致则说明地址有效，返回true
	if bytes.Compare(checkSumBytes,checkBytes) == 0 {
		return true
	}

	return false
}

//获取用户地址（公钥经过Ripemd160Hash,再经过base58编码,最后生成地址）
func GetAddress(pk string) []byte  {

	pkey, err := hex.DecodeString(pk)//对数据库中的公钥进行解码，前提是公钥是经过EncodeString存入的。
	if err != nil {
		return nil
	}
	//调用Ripemd160Hash返回160位的Pub Key hash
	ripemd160Hash := Ripemd160Hash(pkey)

	//将version+Pub Key hash
	version_ripemd160Hash := append([]byte{version},ripemd160Hash...)

	//调用CheckSum方法返回前四个字节的checksum
	checkSumBytes := CheckSum(version_ripemd160Hash)

	//将version+Pub Key hash+ checksum生成25个字节
	bytes := append(version_ripemd160Hash,checkSumBytes...)

	//将这25个字节进行base58编码并返回
	return Base58Encode(bytes)
}
//取前4个字节
func CheckSum(payload []byte) []byte {
	//这里传入的payload其实是version+Pub Key hash，对其进行两次256运算
	hash1 := sha256.Sum256(payload)

	hash2 := sha256.Sum256(hash1[:])

	return hash2[:addressChecksumLen] //返回前四个字节，为CheckSum值
}

//对公钥进行哈希256运算返回160位的哈希值
func Ripemd160Hash(publicKey []byte) []byte {

	//将传入的公钥进行256运算，返回256位hash值
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hash := hash256.Sum(nil)

	//将上面的256位hash值进行160运算，返回160位的hash值
	ripemd160 := ripemd160.New()
	ripemd160.Write(hash)

	return ripemd160.Sum(nil) //返回Pub Key hash
}


// 创建钱包
//func NewWallet() *Wallet {
//
//	privateKey,publicKey := newKeyPair()
//
//	return &Wallet{privateKey,publicKey}
//}


// 通过私钥产生公钥
//func newKeyPair() (ecdsa.PrivateKey,[]byte) {
//	//这是一个曲线对象
//	curve := elliptic.P256()
//	//通过椭圆曲线加密算法生成私钥
//	private, err := ecdsa.GenerateKey(curve, rand.Reader)
//	if err != nil {
//		log.Panic(err)
//	}
//	//由私钥生成公钥
//	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
//
//	return *private, pubKey
//}
//
//
//func main() {
//	//wallet := NewWallet()
//
//	address := wallet.GetAddress()
//
//	fmt.Printf("address：%s\n",address)
//
//	isValid := IsValidForAdress(address)
//
//	fmt.Printf("%s 这个地址为 %v\n",address,isValid)
//}
