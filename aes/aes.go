package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AES/ECB/PKCS7模式解密--签名解密方式
func AesECBDecrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	ecb := NewECBDecryptEr(block)
	retData := make([]byte, len(data))
	ecb.CryptBlocks(retData, data)
	// 解PKCS7填充
	retData = PKCS7UnPadding(retData)
	return retData, nil
}

// AES/ECB/PKCS7模式加密--签名加密方式
func AesECBEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	ecb := NewECBEncryptEr(block)
	// 加PKCS7填充
	content := PKCS7Padding(data, block.BlockSize())
	encryptData := make([]byte, len(content))
	// 生成加密数据
	ecb.CryptBlocks(encryptData, content)
	return encryptData, nil
}

// PKCS7解填充/和PKCS5填充一样,只是填充字段多少的区别
func PKCS7UnPadding(encrypt []byte) []byte {
	length := len(encrypt)
	unPadding := int(encrypt[length-1])
	return encrypt[:(length - unPadding)]
}

// PKCS7加填充/和PKCS5填充一样,只是填充字段多少的区别
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// ----------------------------------------------------------------------------------------------
// ecb加密方法
type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncryptEr ecb

func NewECBEncryptEr(b cipher.Block) cipher.BlockMode {
	return (*ecbEncryptEr)(newECB(b))
}

func (x *ecbEncryptEr) BlockSize() int { return x.blockSize }

func (x *ecbEncryptEr) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// ecb解密方法
type ecbDecryptEr ecb

func NewECBDecryptEr(b cipher.Block) cipher.BlockMode {
	return (*ecbDecryptEr)(newECB(b))
}

func (x *ecbDecryptEr) BlockSize() int { return x.blockSize }

func (x *ecbDecryptEr) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
