package iyzilink

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type CreateLinkRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                float64 `json:"price"`
	CurrencyCode         string  `json:"currencyCode"`
	AddressIgnorable     bool    `json:"addressIgnorable"`
	StockEnabled         bool    `json:"stockEnabled"`
	StockCount           int64   `json:"stockCount"`
	InstallmentRequested bool    `json:"installmentRequested"`
	EncodedImageFile     string  `json:"encodedImageFile"`
}
type CreateLinkResponse struct {
	Status         string `json:"status"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationID string `json:"conversationId"`
	Data           struct {
		Token    string `json:"token"`
		URL      string `json:"url"`
		ImageURL string `json:"imageUrl"`
	} `json:"data"`
}
type DeleteLinkResponse struct {
	Status     string `json:"status"`
	SystemTime int64  `json:"systemTime"`
}
type LinkDetailResponse struct {
	Status         string `json:"status"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationID string `json:"conversationId"`
	Data           struct {
		Name                 string        `json:"name"`
		ConversationID       string        `json:"conversationId"`
		Description          string        `json:"description"`
		Price                float64       `json:"price"`
		CurrencyID           int           `json:"currencyId"`
		CurrencyCode         string        `json:"currencyCode"`
		Token                string        `json:"token"`
		ProductType          string        `json:"productType"`
		ProductStatus        string        `json:"productStatus"`
		MerchantID           int           `json:"merchantId"`
		URL                  string        `json:"url"`
		ImageURL             string        `json:"imageUrl"`
		AddressIgnorable     bool          `json:"addressIgnorable"`
		SoldCount            int           `json:"soldCount"`
		InstallmentRequested bool          `json:"installmentRequested"`
		StockEnabled         bool          `json:"stockEnabled"`
		StockCount           int           `json:"stockCount"`
		PresetPriceValues    []interface{} `json:"presetPriceValues"`
		FlexibleLink         bool          `json:"flexibleLink"`
	} `json:"data"`
}
type IyziOptions struct {
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
	BaseUrl   string `json:"baseUrl"`
}

func CreateLink(c CreateLinkRequest, options IyziOptions) (CreateLinkResponse, error) {

	var res CreateLinkResponse
	authStr, err := createAuthStr(options.ApiKey, options.SecretKey, options.BaseUrl, c)
	if err != nil {
		return CreateLinkResponse{}, err
	}
	client := resty.New()
	ma, err := json.Marshal(c)
	resp, err := client.R().
		SetHeader("Authorization", authStr).
		SetHeader("Content-Type", "application/json").
		SetHeader("Content-Length", fmt.Sprintf(`"%s"`, strconv.Itoa(len(ma)))).
		SetHeader("Cache-Control", "no-cache").
		SetBody(c).
		SetResult(&res).
		Post(options.BaseUrl)

	if err != nil {
		return res, err
	}

	if resp.IsError() {
		return res, fmt.Errorf(string(resp.Body()))
	}

	if res.Status != "success" {
		return res, fmt.Errorf(string(resp.Body()))
	}
	return res, nil

}
func GetLinkDetail(token string, options IyziOptions) (LinkDetailResponse, error) {
	var res LinkDetailResponse
	endPoint := fmt.Sprintf("%s/%s", options.BaseUrl, token)
	authStr, err := createAuthStr(options.ApiKey, options.SecretKey, endPoint, nil)
	if err != nil {
		return LinkDetailResponse{}, err
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", authStr).
		SetHeader("Content-Type", "application/json").
		SetResult(&res).
		Get(endPoint)
	if err != nil {
		return res, err
	}
	fmt.Println("Response Body:", string(resp.Body()))
	if resp.IsError() {
		return res, fmt.Errorf(string(resp.Body()))
	}

	if res.Status != "success" {
		return res, fmt.Errorf(string(resp.Body()))
	}
	return res, nil
}
func DeleteLink(token string, options IyziOptions) (DeleteLinkResponse, error) {
	var res DeleteLinkResponse
	endPoint := fmt.Sprintf("%s/%s", options.BaseUrl, token)
	authStr, err := createAuthStr(options.ApiKey, options.SecretKey, endPoint, nil)
	if err != nil {
		return DeleteLinkResponse{}, err
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", authStr).
		SetHeader("Content-Type", "application/json").
		SetResult(&res).
		Delete(endPoint)
	if err != nil {
		return res, err
	}
	fmt.Println("Response Body:", string(resp.Body()))
	if resp.IsError() {
		return res, fmt.Errorf(string(resp.Body()))
	}

	if res.Status != "success" {
		return res, fmt.Errorf(string(resp.Body()))
	}
	return res, nil
}
