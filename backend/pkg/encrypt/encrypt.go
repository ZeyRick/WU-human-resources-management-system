package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func AesCbcEncrypt(src string, key []byte) (string, error) {
	initialVector := "03f6b349a565fcdc"
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted), nil
}


func AesCbcDecrypt(crypt []byte, key []byte) (string, error) {
	initialVector := "03f6b349a565fcdc"
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(crypt) == 0 {
		return "", errors.New("NO CONTENT")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)
	return string(PKCS5Trimming(decrypted)), nil
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
		concatArray = append(concatArray, k+"md5Key="+ToString(payload[k]))
	}

	concatMd5String := strings.Join(concatArray[:], "-") + key
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
