package wxpay

import (
	"encoding/xml"
)

// PlaceOrderResult represent place order reponse message from weixin pay.
// For field explanation refer to: http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_1
type PlaceOrderResult struct {
	XMLName     xml.Name `xml:"xml"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`
	AppId       string   `xml:"appid"`
	MchId       string   `xml:"mch_id"`
	DeviceInfo  string   `xml:"device_info"`
	NonceStr    string   `xml:"nonce_str"`
	Sign        string   `xml:"sign"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDesc string   `xml:"err_code_des"`
	TradeType   string   `xml:"trade_type"`
	PrepayId    string   `xml:"prepay_id"`
	CodeUrl     string   `xml:"code_url"`
	MwebUrl     string   `xml:"mweb_url"`
}

func (this *PlaceOrderResult) ToMap() map[string]string {
	retMap, err := ToMap(this)
	if err != nil {
		panic(err)
	}

	return retMap
}

// Parse the reponse message from weixin pay to struct of PlaceOrderResult
func ParsePlaceOrderResult(resp []byte) (PlaceOrderResult, error) {
	placeOrderResult := PlaceOrderResult{}
	err := xml.Unmarshal(resp, &placeOrderResult)
	if err != nil {
		return placeOrderResult, err
	}

	return placeOrderResult, nil
}

// QueryOrder Result represent query response message from weixin pay
// Refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_2&index=4
type QueryOrderResult struct {
	XMLName        xml.Name `xml:"xml"`
	ReturnCode     string   `xml:"return_code"`
	ReturnMsg      string   `xml:"return_msg"`
	AppId          string   `xml:"appid"`
	MchId          string   `xml:"mch_id"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	ResultCode     string   `xml:"result_code"`
	ErrCode        string   `xml:"err_code"`
	ErrCodeDesc    string   `xml:"err_code_des"`
	DeviceInfo     string   `xml:"device_info"`
	OpenId         string   `xml:"open_id"`
	IsSubscribe    string   `xml:"is_subscribe"`
	TradeType      string   `xml:"trade_type"`
	TradeState     string   `xml:"trade_state"`
	TradeStateDesc string   `xml:"trade_state_desc"`
	BankType       string   `xml:"bank_type"`
	TotalFee       string   `xml:"total_fee"`
	FeeType        string   `xml:"fee_type"`
	CashFee        string   `xml:"cash_fee"`
	CashFeeType    string   `xml:"cash_fee_type"`
	CouponFee      string   `xml:"coupon_fee"`
	CouponCount    string   `xml:"coupon_count"`
	TransactionId  string   `xml:"transaction_id"`
	OrderId        string   `xml:"out_trade_no"`
	Attach         string   `xml:"attach"`
	TimeEnd        string   `xml:"time_end"`
}

func (this *QueryOrderResult) ToMap() map[string]string {
	retMap, err := ToMap(this)
	if err != nil {
		panic(err)
	}

	return retMap
}

func ParseQueryOrderResult(resp []byte) (QueryOrderResult, error) {
	queryOrderResult := QueryOrderResult{}
	err := xml.Unmarshal(resp, &queryOrderResult)
	if err != nil {
		return queryOrderResult, err
	}

	return queryOrderResult, nil
}

//微信支付异步通知服务器结果
type NoticeOrderResult struct {
	XMLName            xml.Name `xml:"xml"`
	ReturnCode         string   `xml:"return_code"`
	ReturnMsg          string   `xml:"return_msg"`
	AppId              string   `xml:"appid"`
	MchId              string   `xml:"mch_id"`
	DeviceInfo         string   `xml:"device_info"`
	NonceStr           string   `xml:"nonce_str"`
	Sign               string   `xml:"sign"`
	SignType           string   `xml:"sign_type"`
	SubMchId           string   `xml:"sub_mch_id"`
	ResultCode         string   `xml:"result_code"`
	ErrCode            string   `xml:"err_code"`
	ErrCodeDesc        string   `xml:"err_code_des"`
	OpenId             string   `xml:"openid"`
	IsSubscribe        string   `xml:"is_subscribe"`
	TradeType          string   `xml:"trade_type"`
	BankType           string   `xml:"bank_type"`
	TotalFee           string   `xml:"total_fee"`
	SettlementTotalFee string   `xml:"settlement_total_fee"`
	FeeType            string   `xml:"fee_type"`
	CashFee            string   `xml:"cash_fee"`
	CashFeeType        string   `xml:"cash_fee_type"`
	CouponFee          string   `xml:"coupon_fee"`
	CouponFee0         string   `xml:"coupon_fee_0"`
	CouponCount        string   `xml:"coupon_count"`
	CouponType         string   `xml:"coupon_type"`
	CouponType0        string   `xml:"coupon_type_0"`
	CouponId           string   `xml:"coupon_id"`
	CouponId0          string   `xml:"coupon_id_0"`
	TransactionId      string   `xml:"transaction_id"`
	OutTradeNo         string   `xml:"out_trade_no"`
	Attach             string   `xml:"attach"`
	TimeEnd            string   `xml:"time_end"`
}

func (this *NoticeOrderResult) ToMap() map[string]string {
	retMap, err := ToMap(this)
	if err != nil {
		panic(err)
	}

	return retMap
}
func ParseNoticeOrderResult(resp []byte) (NoticeOrderResult, error) {
	noticeOrderResult := NoticeOrderResult{}
	err := xml.Unmarshal(resp, &noticeOrderResult)
	if err != nil {
		return noticeOrderResult, err
	}

	return noticeOrderResult, nil
}
