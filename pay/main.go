package main

import "fmt"

func main(){
	// 这里切记 字典类型的数据是需要初始化的
	phone := &Phone{
		PayMap:make(map[string]Pay,10),
	}

	// 这里是用于开通自己有哪些支付方式
	//phone.OpenWeChatPay()
	//phone.OpenAliPay()
	//phone.OpenPay("weChatPay",&WeChatPay{})
	phone.OpenPay("aLiPay",&AliPay{})
	err := phone.PayMoney("weChatPay",100)
	if err != nil{
		// 如果微信支付失败了，用支付宝支付
		fmt.Printf("支付失败，失败原因：%v\n",err)
		fmt.Println("使用支付宝支付")
		err = phone.PayMoney("aLiPay",100)
		if err != nil{
			fmt.Printf("支付失败，失败原因：%v\n",err)
			return
		}
	}
	fmt.Println("支付成功，欢迎再次光临")
}
