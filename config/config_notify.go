package config

// 定义一个通知的接口
type Notifyer interface {
	Callback(*Config)
}
