package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func AesCbcEncrypt(src string, key []byte) ([]byte, error) {
	initialVector := "1010101010101010"
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if src == "" {
		return nil, errors.New("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted, nil
}

func AesCbcDecrypt(crypt []byte, key []byte) ([]byte, error) {
	initialVector := "1010101010101010"
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(crypt) == 0 {
		return nil, errors.New("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)
	return PKCS5Trimming(decrypted), nil
}
func hashMd5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenerateSignature(payload map[string]interface{}, key string) string {
	keys := []string{}
	for k := range payload {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	concatArray := []string{}
	for _, k := range keys {
		concatArray = append(concatArray, k+"="+ToString(payload[k]))
	}

	concatMd5String := strings.Join(concatArray[:], "&") + key
	return hashMd5(concatMd5String)
}

func VerifySignature(sign string, payload map[string]interface{}, key string) bool {
	return sign == GenerateSignature(payload, key)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func ToString(s any) string {
	if s == "" || s == nil {
		return ""
	}
	switch reflect.TypeOf(s).Kind() {
	case reflect.String:
		return fmt.Sprintf("%v", s)
	case reflect.Int:
		return fmt.Sprintf("%v", s)
	case reflect.Float64:
		return fmt.Sprintf("%v", s)
	case reflect.Bool:
		return strconv.FormatBool(s.(bool))
	default:
		return s.(string)
	}
}
