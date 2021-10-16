package utils

import (
	"crypto/aes"
)

const eapiKey = "e82ckenh8dichen8"
const cacheKey = ")(13daqP@ssw0rd~"
const markerKey = "#14ljk_!\\]&0U<'("

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func encryptECB(data, keyStr string) (encrypted []byte) {
	origData := []byte(data)
	key := []byte(keyStr)
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func decryptECB(encrypted []byte, keyStr string) (decrypted []byte) {
	key := []byte(keyStr)
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// EapiEncrypt eapi 加密
func EapiEncrypt(data string) (encrypted []byte) {
	return encryptECB(data, eapiKey)
}

// MarkerEncrypt 163 key 加密
func MarkerEncrypt(data string) (encrypted []byte) {
	return encryptECB(data, markerKey)
}

// CacheKeyEncrypt cache_key 加密
func CacheKeyEncrypt(data string) (encrypted []byte) {
	return encryptECB(data, cacheKey)
}

// EapiDecrypt eapi 解密
func EapiDecrypt(encrypted []byte) (decrypted []byte) {
	return decryptECB(encrypted, eapiKey)
}

// MarkerDecrypt 163 key 解密
func MarkerDecrypt(encrypted []byte) (decrypted []byte) {
	return decryptECB(encrypted, markerKey)
}
