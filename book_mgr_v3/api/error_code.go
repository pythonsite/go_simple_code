package main



const(
	ErrSuccess = 0
	ErrInvalidParameter = 1001
	ErrServerBusy = 1002

)

func getMessage(code int)(msg string){
	switch code{
	case ErrSuccess:
		msg = "success"
	case ErrInvalidParameter:
		msg = "invalid parameter"
	case ErrServerBusy:
		msg = "server busy"
	default:
		msg = "unknown error"
	}
	return
}
