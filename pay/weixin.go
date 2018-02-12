package main

import "fmt"

type WeChatPay struct {

}

// 这里也是实现了Pay接口中的pay方法
func (w *WeChatPay) pay(userId int64,money float32) error{
	fmt.Println("1.连接到微信支付的服务器")
	fmt.Println("2.连接到对应的用户")
	fmt.Println("3.检查余额")
	fmt.Println("4.扣钱")
	fmt.Println("5.返回支付是否成功")

	return nil
}