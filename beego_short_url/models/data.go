package models


type Long2ShortRequest struct {
	OriginUrl string `json:"origin_url"`
}

type ResponseHeader struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Long2ShortResponse struct {
	ResponseHeader
	ShortUrl string `json:"short_url"`
}

type Short2LongRequest struct {
	ShortUrl string `json:"short_url"`
}

type Short2LongResponse struct {
	ResponseHeader
	OriginUrl string `json:"origin_url"`
}

type ShortUrl struct {
	ShortUrl string `json:"short_url" db:"short_url"`
}


