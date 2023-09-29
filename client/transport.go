package client

import "net/http"

type WildberriesAuthTransport struct {
	transport http.RoundTripper
	token     string
}

func (t *WildberriesAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", t.token)
	return t.transport.RoundTrip(req)
}

func NewWildberriesAuthTransport(token string) *WildberriesAuthTransport {
	return &WildberriesAuthTransport{transport: http.DefaultTransport, token: token}
}
