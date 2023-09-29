package go_wildberries

type WildberriesResponse[T any] struct {
	Data             T        `json:"data"`
	Error            bool     `json:"error"`
	ErrorText        string   `json:"errorText"`
	AdditionalErrors []string `json:"additionalErrors"`
}
