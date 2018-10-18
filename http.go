package wxpay

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WxPay is abstact of Transaction handler. With WxPay, we can get prepay id
type WxPay struct {
	Config *WxConfig
}

// Initialized the WxPay with specific config
func NewWxPay(cfg *WxConfig) (*WxPay, error) {
	if cfg.AppId == "" ||
		cfg.MchId == "" ||
		cfg.AppKey == "" ||
		cfg.NotifyUrl == "" ||
		cfg.QueryOrderUrl == "" ||
		cfg.PlaceOrderUrl == "" {
		return &WxPay{Config: cfg}, errors.New("config field canot empty string")
	}

	return &WxPay{Config: cfg}, nil
}

// Submit the order to weixin pay and return the prepay id if success,
// Prepay id is used for app to start a payment
// If fail, error is not nil, check error for more information
func (this *WxPay) AppSubmit(orderId string, amount int, desc string, attach string, clientIp string) (string, error) {

	odrInXml := this.signedOrderRequestXmlString("APP", orderId, fmt.Sprint(amount), desc, attach, clientIp, "")
	resp, err := doHttpPost(this.Config.PlaceOrderUrl, []byte(odrInXml))
	if err != nil {
		return "", err
	}

	placeOrderResult, err := ParsePlaceOrderResult(resp)
	if err != nil {
		return "", err
	}

	//Verify the sign of response
	resultInMap := placeOrderResult.ToMap()

	if placeOrderResult.ReturnCode != "SUCCESS" {
		return "", fmt.Errorf("return code:%s, return desc:%s", placeOrderResult.ReturnCode, placeOrderResult.ReturnMsg)
	}

	if placeOrderResult.ResultCode != "SUCCESS" {
		return "", fmt.Errorf("resutl code:%s, result desc:%s", placeOrderResult.ErrCode, placeOrderResult.ErrCodeDesc)
	}
	wantSign := Sign(resultInMap, this.Config.AppKey)
	gotSign := resultInMap["sign"]
	if wantSign != gotSign {
		return "", fmt.Errorf("sign not match, want:%s, got:%s", wantSign, gotSign)
	}

	return placeOrderResult.PrepayId, nil
}

func (this *WxPay) H5Submit(orderId string, amount int, desc string, attach string, clientIp, H5SceneInfo string) (string, error) {

	odrInXml := this.signedOrderRequestXmlString("MWEB", orderId, fmt.Sprint(amount), desc, attach, clientIp, H5SceneInfo)
	resp, err := doHttpPost(this.Config.PlaceOrderUrl, []byte(odrInXml))
	if err != nil {
		return "", err
	}

	placeOrderResult, err := ParsePlaceOrderResult(resp)
	if err != nil {
		return "", err
	}

	//Verify the sign of response
	resultInMap := placeOrderResult.ToMap()

	if placeOrderResult.ReturnCode != "SUCCESS" {
		return "", fmt.Errorf("return code:%s, return desc:%s", placeOrderResult.ReturnCode, placeOrderResult.ReturnMsg)
	}

	if placeOrderResult.ResultCode != "SUCCESS" {
		return "", fmt.Errorf("resutl code:%s, result desc:%s", placeOrderResult.ErrCode, placeOrderResult.ErrCodeDesc)
	}
	wantSign := Sign(resultInMap, this.Config.AppKey)
	gotSign := resultInMap["sign"]
	if wantSign != gotSign {
		return "", fmt.Errorf("sign not match, want:%s, got:%s", wantSign, gotSign)
	}

	return placeOrderResult.PrepayId, nil
}

func (this *WxPay) newQueryXml(out_trade_no string) string {
	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["mch_id"] = this.Config.MchId
	param["out_trade_no"] = out_trade_no
	param["nonce_str"] = NewNonceString()

	sign := Sign(param, this.Config.AppKey)
	param["sign"] = sign

	return ToXmlString(param)
}

// Query the order from weixin pay server by transaction id of weixin pay
func (this *WxPay) Query(transId string) (QueryOrderResult, error) {
	queryOrderResult := QueryOrderResult{}

	queryXml := this.newQueryXml(transId)
	// fmt.Println(queryXml)
	resp, err := doHttpPost(this.Config.QueryOrderUrl, []byte(queryXml))
	if err != nil {
		return queryOrderResult, nil
	}

	queryOrderResult, err = ParseQueryOrderResult(resp)
	if err != nil {
		return queryOrderResult, err
	}

	//verity sign of response
	resultInMap := queryOrderResult.ToMap()
	wantSign := Sign(resultInMap, this.Config.AppKey)
	gotSign := resultInMap["sign"]
	if wantSign != gotSign {
		return queryOrderResult, fmt.Errorf("sign not match, want:%s, got:%s", wantSign, gotSign)
	}

	return queryOrderResult, nil
}

// NewPaymentRequest build the payment request structure for app to start a payment.
// Return stuct of PaymentRequest, please refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_12&index=2
func (this *WxPay) NewPaymentRequest(prepayId string) PaymentRequest {
	noncestr := NewNonceString()
	timestamp := NewTimestampString()

	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["partnerid"] = this.Config.MchId
	param["prepayid"] = prepayId
	param["package"] = "Sign=WXPay"
	param["noncestr"] = noncestr
	param["timestamp"] = timestamp

	sign := Sign(param, this.Config.AppKey)

	payRequest := PaymentRequest{
		AppId:     this.Config.AppId,
		PartnerId: this.Config.MchId,
		PrepayId:  prepayId,
		Package:   "Sign=WXPay",
		NonceStr:  noncestr,
		Timestamp: timestamp,
		Sign:      sign,
	}

	return payRequest
}

func (this *WxPay) newOrderRequest(payType, orderId, amount, desc, attach, clientIp, H5SceneInfo string) map[string]string {
	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["attach"] = attach //optional
	param["body"] = desc
	param["mch_id"] = this.Config.MchId
	param["nonce_str"] = NewNonceString()
	param["notify_url"] = this.Config.NotifyUrl
	param["out_trade_no"] = orderId
	param["spbill_create_ip"] = clientIp
	param["total_fee"] = amount
	param["trade_type"] = payType

	if payType == "MWEB" {
		param["scene_info"] = H5SceneInfo
	}

	return param
}

func (this *WxPay) signedOrderRequestXmlString(payType, orderId, amount, desc, attach, clientIp, H5SceneInfo string) string {
	order := this.newOrderRequest(payType, orderId, amount, desc, attach, clientIp, H5SceneInfo)
	sign := Sign(order, this.Config.AppKey)
	// fmt.Println(sign)

	order["sign"] = sign

	return ToXmlString(order)
}

// doRequest post the order in xml format with a sign
func doHttpPost(targetUrl string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respData, nil
}
