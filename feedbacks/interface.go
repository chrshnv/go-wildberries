package feedbacks

import (
	"encoding/json"
	"fmt"
	go_wildberries "github.com/chrshnv/go-wildberries"
	"io"
	"net/http"
)

type WildberriesFeedbacksAPI interface {
	Fetch(isAnswered bool, take, skip int) (go_wildberries.WildberriesResponse[WildberriesFeedbackListResponse], error)
}

type wildberriesFeedbacksAPI struct {
	client http.Client
}

func (w *wildberriesFeedbacksAPI) Fetch(isAnswered bool, take, skip int) (go_wildberries.WildberriesResponse[WildberriesFeedbackListResponse], error) {
	uri := fmt.Sprintf(
		"%s?isAnswered=%t&take=%d&skip=%d",
		"https://feedbacks-api.wildberries.ru/api/v1/feedbacks",
		isAnswered,
		take,
		skip,
	)

	resp, err := w.client.Get(uri)
	if err != nil {
		return go_wildberries.WildberriesResponse[WildberriesFeedbackListResponse]{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return go_wildberries.WildberriesResponse[WildberriesFeedbackListResponse]{}, err
	}

	var result go_wildberries.WildberriesResponse[WildberriesFeedbackListResponse]
	err = json.Unmarshal(body, &result)
	if err != nil {
		return go_wildberries.WildberriesResponse[WildberriesFeedbackListResponse]{}, err
	}

	return result, nil
}

func NewFeedbacksAPI(token string) WildberriesFeedbacksAPI {
	return &wildberriesFeedbacksAPI{client: http.Client{Transport: go_wildberries.NewWildberriesAuthTransport(token)}}
}
