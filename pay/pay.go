package main


type Pay interface {
	pay(userId int64,money float32) error
}
