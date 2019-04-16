package algorithm

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"util"
	"io/ioutil"
	"fmt"
	"math/big"
	"os"
	"io"
	"crypto/sha256"
)

//获取公钥和私钥
func GetKey() (*ecdsa.PrivateKey,*ecdsa.PublicKey,error){
	//生成一对ECDSA公钥和私钥
	prk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	util.CheckErr(err)
	pub := &prk.PublicKey
	//fmt.Println("私钥 %s",prk)
	//fmt.Println("公钥 %s",pub)
	return prk,pub,err
}
//非对称加密
//公钥加密
func ECCEncrypt(pt []byte, pub *ecdsa.PublicKey) ([]byte, error) {
	pub2 := ImportECDSAPublic(pub)//将ECDSA公钥转换为ECIES公钥
	ct, err := Encrypt(rand.Reader, pub2, pt, nil, nil)
	return ct, err
}

//私钥解密
func ECCDecrypt(ct []byte, prk *ecdsa.PrivateKey) ([]byte, error) {
	prk2 := ImportECDSA(prk)//将ECDSA私钥转换为ECIES私钥
	pt, err := prk2.Decrypt(ct, nil, nil)
	return pt, err
}

func toECDSA(d []byte, strict bool) (*ecdsa.PrivateKey,error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = elliptic.P256()
	if strict && 8*len(d) != priv.Params().BitSize {
		return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)
	if priv.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)

	return priv,nil
}

// 私钥 -> []byte
func PrivateKeyToByte(priv *ecdsa.PrivateKey) []byte {
	if priv == nil {
		return nil
	}
	prb1:=PaddedBigBytes(priv.D, priv.Params().BitSize/8)
	return prb1
}

// ToECDSA creates a private key with the given D value.
// []byte -> 私钥
// ToECDSA creates a private key with the given D value.
func ToECDSA(d []byte) (*ecdsa.PrivateKey,error) {
	return toECDSA(d, true)
}

// 公钥 -> []byte
func PublicKeyToByte(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	pub1:=elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	//fmt.Println("公钥转换后 %s",ToECDSAPub(pub1))
	return pub1
}

// []byte -> 公钥
func ToECDSAPub(pub []byte) *ecdsa.PublicKey {
	if len(pub) == 0 {
		return nil
	}
	x, y := elliptic.Unmarshal(elliptic.P256(), pub)
	turnpuk:=&ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}
	//fmt.Println("公钥转换后 %s",turnpuk)
	return turnpuk
}

// 将secp256k1私钥写入文件，保存为十六进制编码
func SavePrivateKey(file string, key *ecdsa.PrivateKey) error {
	k := hex.EncodeToString(PrivateKeyToByte(key))
	//turnprk,err:=ToECDSA([]byte(k))
	//util.CheckErr(err)
	//fmt.Println("私钥转换后 %s",turnprk,err)
	return ioutil.WriteFile(file, []byte(k), 0600)
}

// 将公钥写入文件，保存为十六进制编码
func SavePublicKey(file string, key *ecdsa.PublicKey) error {
	k := hex.EncodeToString(PublicKeyToByte(key))
	return ioutil.WriteFile(file, []byte(k), 0600)
}

//把私钥从文件中读取出来并转化成私钥格式
func LoadECDSA(file string) (*ecdsa.PrivateKey, error) {
	buf := make([]byte, 64)
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	if _, err := io.ReadFull(fd, buf); err != nil {
		return nil, err
	}
	//私钥从16进制解码
	key, err := hex.DecodeString(string(buf))
	if err != nil {
		return nil, err
	}
	return ToECDSA(key)
}

//用私钥进行签名
func eccSign(data []byte, path string) ([]byte, error) {
	//取得私钥
	privateKey,err := LoadECDSA(path)
	//计算交易信息的哈希值，已经哈希过就不用再做这一操作
	hash := sha256.New()
	//填入数据
	hash.Write(data)
	bytes := hash.Sum(nil)

	//对哈希值生成数字签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, bytes)
	if err != nil {
		return nil, err
	}
	params := privateKey.Curve.Params()
	curveOrderByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveOrderByteSize*2)
	copy(signature[curveOrderByteSize-len(rBytes):], rBytes)
	copy(signature[curveOrderByteSize*2-len(sBytes):], sBytes)
	return signature, nil
}

func eccVerify(data, signature []byte, pk []byte) bool {
	////读取公钥pk
	//从数据库中读取公钥pk的值后解码
	key, err := hex.DecodeString(string(pk))
	if err != nil {
		return false
	}
	//调用ToECDSAPub函数，转换为公钥格式
	publicKey:=ToECDSAPub(key)
	curveOrderByteSize := publicKey.Curve.Params().P.BitLen() / 8
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(signature[:curveOrderByteSize])
	s.SetBytes(signature[curveOrderByteSize:])

	return ecdsa.Verify(publicKey, data, r, s)
}