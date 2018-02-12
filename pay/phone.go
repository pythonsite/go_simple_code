package main

import "fmt"

type Phone struct {
	PayMap map[string]Pay
}

func (p *Phone) OpenWeChatPay(){
	weChatPay := &WeChatPay{}
	p.PayMap["weChatPay"] = weChatPay
}

func (p *Phone) OpenAliPay(){
	AliPay := &AliPay{}
	p.PayMap["aLiPay"] = AliPay
}

func (p *Phone) PayMoney(name string,money float32)(err error){
	pay,ok:= p.PayMap[name]
	if !ok{
		err = fmt.Errorf("不支持【%s】支付方式",name)
		return
	}
	err = pay.pay(1024,money)
	return
}
