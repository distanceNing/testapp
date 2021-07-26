package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func VerifySignBySha1(data string, sign string) bool {
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(data))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	encode_bs := hex.EncodeToString(bs)
	return encode_bs == sign;
}

// var appId = "wx4f4bc4dec97d474b"
//var sessionKey = "tiihtNczf5v6AKRyjwEUhQ=="
//var encryptedData = "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
//var iv = "r7BXXKkLb8qrSNn05n0qiA=="
func AesDecrypt(key string, data string, iv string) []byte {

	decode_key, _ := base64.StdEncoding.DecodeString(key)
	decode_iv, _ := base64.StdEncoding.DecodeString(iv)
	block, err := aes.NewCipher(decode_key)
	if err != nil {
		return make([]byte, 0)
	}
	blockMode := cipher.NewCBCDecrypter(block, decode_iv)

	// crypted := []byte(encryptedData)
	decode_data, _ := base64.StdEncoding.DecodeString(data)
	origData := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origData, decode_data)
	return ZeroUnPadding(origData)
}
