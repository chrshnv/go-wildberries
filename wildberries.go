package go_wildberries

import "net/http"

type WildberriesAPI struct {
	client http.Client
}

func NewWildberriesAPI(auth string) *WildberriesAPI {
	return &WildberriesAPI{client: http.Client{Transport: NewWildberriesAuthTransport(auth)}}
}
