package feedbacks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chrshnv/go-wildberries/client"
	"github.com/chrshnv/go-wildberries/responses"
)

type WildberriesFeedbacksAPI interface {
	Fetch(isAnswered bool, take, skip int) (responses.WildberriesResponse[WildberriesFeedbackListResponse], error)
	WorkWithFeedbackAnswer(id string, text string) error
	WorkWithFeedbackView(id string, wasViewed bool) error
}

type wildberriesFeedbacksAPI struct {
	client http.Client
}

func (w *wildberriesFeedbacksAPI) Fetch(isAnswered bool, take, skip int) (responses.WildberriesResponse[WildberriesFeedbackListResponse], error) {
	uri := fmt.Sprintf(
		"%s?isAnswered=%t&take=%d&skip=%d",
		"https://feedbacks-api.wildberries.ru/api/v1/feedbacks",
		isAnswered,
		take,
		skip,
	)

	resp, err := w.client.Get(uri)
	if err != nil {
		return responses.WildberriesResponse[WildberriesFeedbackListResponse]{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return responses.WildberriesResponse[WildberriesFeedbackListResponse]{}, err
	}

	var result responses.WildberriesResponse[WildberriesFeedbackListResponse]
	err = json.Unmarshal(body, &result)
	if err != nil {
		return responses.WildberriesResponse[WildberriesFeedbackListResponse]{}, err
	}

	return result, nil
}

func (w *wildberriesFeedbacksAPI) WorkWithFeedbackAnswer(id string, text string) error {
	body := WildberriesFeedbackPatchAnswer{
		Id:   id,
		Text: text,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, "https://feedbacks-api.wb.ru/api/v1/feedbacks", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	_, err = w.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (w *wildberriesFeedbacksAPI) WorkWithFeedbackView(id string, wasViewed bool) error {
	body := WildberriesFeedbackPatchViewed{
		Id:        id,
		WasViewed: wasViewed,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, "https://feedbacks-api.wb.ru/api/v1/feedbacks", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	_, err = w.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func NewFeedbacksAPI(token string) WildberriesFeedbacksAPI {
	return &wildberriesFeedbacksAPI{client: http.Client{Transport: client.NewWildberriesAuthTransport(token)}}
}
