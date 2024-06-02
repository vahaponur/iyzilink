package iyzilink

import (
	"encoding/json"
	"fmt"
	"testing"
)

var opt = IyziOptions{
	ApiKey:    "YOUR_API_KEY",
	SecretKey: "YOUR_SECRET_KEY",
	BaseUrl:   "https://api.iyzipay.com/v2/iyzilink/products",
}

func TestCreateLink(t *testing.T) {
	c := CreateLinkRequest{

		Name:                 "Davetiye 100 Adet",
		Price:                1560,
		StockCount:           1,
		StockEnabled:         true,
		AddressIgnorable:     false,
		InstallmentRequested: true,
		CurrencyCode:         "TRY",
		EncodedImageFile:     "base64Img",
		Description:          "Resim Temsilidir",
	}
	res, err := CreateLink(c, opt)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
func TestGetLinkDetail(t *testing.T) {
	res, err := GetLinkDetail("token", opt)
	if err != nil {
		fmt.Println(err.Error())
	}
	byteRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(byteRes))
}
func TestDeleteLink(t *testing.T) {
	res, err := DeleteLink("token", opt)
	if err != nil {
		fmt.Println(err.Error())
	}
	byteRes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(byteRes))
}
