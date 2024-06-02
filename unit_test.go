package iyzilink

import (
	"fmt"
	"testing"
)

func TestCreateLink(t *testing.T) {
	opt := IyziOptions{
		ApiKey:    "YOUR_API_KEY",
		SecretKey: "YOUR_SECRET_KEY",
		BaseUrl:   "https://api.iyzipay.com/v2/iyzilink/products",
	}
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
	res, err := CreateLink(opt, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
