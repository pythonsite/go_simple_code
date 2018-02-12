package main

// 定义一个支付的接口
type Pay interface {
	pay(userId int64,money float32) error
}
