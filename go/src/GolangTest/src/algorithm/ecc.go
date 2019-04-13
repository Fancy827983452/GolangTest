package algorithm

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"util"
	"io/ioutil"
)

//获取公钥和私钥
func GetKey() (*ecdsa.PrivateKey,*ecdsa.PublicKey,error){
	//生成一对ECDSA公钥和私钥
	prk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	util.CheckErr(err)
	pub := &prk.PublicKey
	return prk,pub,err
}

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

// 私钥 -> []byte
func PrivateKeyToByte(priv *ecdsa.PrivateKey) []byte {
	if priv == nil {
		return nil
	}
	return PaddedBigBytes(priv.D, priv.Params().BitSize/8)
}

// 公钥 -> []byte
func PublicKeyToByte(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(pub.Curve, pub.X, pub.Y)
}

// 将secp256k1私钥写入文件，保存为十六进制编码
func SavePrivateKey(file string, key *ecdsa.PrivateKey) error {
	k := hex.EncodeToString(PrivateKeyToByte(key))
	return ioutil.WriteFile(file, []byte(k), 0600)
}

// 将公钥写入文件，保存为十六进制编码
func SavePublicKey(file string, key *ecdsa.PublicKey) error {
	k := hex.EncodeToString(PublicKeyToByte(key))
	return ioutil.WriteFile(file, []byte(k), 0600)
}