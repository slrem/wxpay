package wxpay

type PaymentRequest struct {
	AppId     string `json:"appId"`
	PartnerId string `json:"partnerId"`
	PrepayId  string `json:"prepayId"`
	Package   string `json:"package"`
	NonceStr  string `json:"nonceStr"`
	Timestamp string `json:"timeStamp"`
	Sign      string `json:"sign"`
}
