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
type IyziOptions struct {
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
	BaseUrl   string `json:"baseUrl"`
}

func CreateLink(options IyziOptions, c CreateLinkRequest) (CreateLinkResponse, error) {

	var res CreateLinkResponse
	authStr := createAuthStr(options.ApiKey, options.SecretKey, options.BaseUrl, c)
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
