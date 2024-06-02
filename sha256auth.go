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

func createSignature(random, requestURL string, data interface{}, secret string) string {

	startIndex := strings.Index(requestURL, "/v2")
	endIndex := strings.Index(requestURL, "?")
	var uriPath string
	if endIndex != -1 {
		uriPath = requestURL[startIndex:endIndex]
	} else {
		uriPath = requestURL[startIndex:]
	}
	jsonData, _ := json.Marshal(data)
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
	return signature
}

func createAuthStr(key, secret string, requestURL string, data interface{}) string {
	randomStr := getRandomString(12)
	signature := createSignature(randomStr, requestURL, data, secret)
	authStr := fmt.Sprintf("apiKey:%s&randomKey:%s&signature:%s", key, randomStr, signature)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authStr))
	fullStr := fmt.Sprintf("IYZWSv2 %s", encodedAuth)

	return fullStr
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
