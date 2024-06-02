# iyzilink
```shell
go get github.com/vahaponur/iyzilink@v1.1.0

```
## CreateLinkRequest

Bu struct, ödeme bağlantısı oluşturmak için gönderilen isteğin yükünü temsil eder.

```go
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
opt := IyziOptions{
ApiKey:    "YOUR_API_KEY",
SecretKey: "YOUR_SECRET_KEY",
BaseUrl:   "https://api.iyzipay.com/v2/iyzilink/products",
}
```
```go
c := CreateLinkRequest{

Name:                 "Davetiye 100 Adet",
Price:                1560,
StockCount:           1,
StockEnabled:         true,
AddressIgnorable:     false,
InstallmentRequested: true,
CurrencyCode:         "TRY",
EncodedImageFile:     "base64image",
Description:          "Resim Temsilidir",
}
res, err := CreateLink(opt, c)
```
```go
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
```
You can reach the link from res.Data.URL
