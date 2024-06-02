package iyzilink

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"unsafe"
)

func createSignature(random, requestURL string, data interface{}, secret string) (string, error) {

	startIndex := strings.Index(requestURL, "/v2")
	endIndex := strings.Index(requestURL, "?")
	var uriPath string
	if endIndex != -1 {
		uriPath = requestURL[startIndex:endIndex]
	} else {
		uriPath = requestURL[startIndex:]
	}
	var jsonData []byte
	var err error
	if data != nil {
		jsonData, err = json.Marshal(data)
		if err != nil {
			return "", err
		}
	}
	var payload string
	if len(jsonData) == 0 {
		payload = uriPath
	} else {
		payload = uriPath + string(jsonData)
	}
	dataToEncrypt := random + payload
	fmt.Println(dataToEncrypt)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(dataToEncrypt))
	signature := hex.EncodeToString(h.Sum(nil))
	return signature, nil
}

func createAuthStr(key, secret string, requestURL string, data interface{}) (string, error) {
	randomStr := getRandomString(12)
	signature, err := createSignature(randomStr, requestURL, data, secret)
	if err != nil {
		return "", err
	}
	authStr := fmt.Sprintf("apiKey:%s&randomKey:%s&signature:%s", key, randomStr, signature)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authStr))
	fullStr := fmt.Sprintf("IYZWSv2 %s", encodedAuth)
	return fullStr, nil
}

func getRandomString(size int) string {
	var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]byte, size)
	rand.Read(s)
	for i := 0; i < size; i++ {
		s[i] = alphabet[s[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&s))
}
