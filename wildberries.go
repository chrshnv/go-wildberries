package go_wildberries

import (
	"github.com/chrshnv/go-wildberries/feedbacks"
)

type WildberriesAPI struct {
	FeedbacksAPI feedbacks.WildberriesFeedbacksAPI
}

func NewWildberriesAPI(auth string) *WildberriesAPI {
	return &WildberriesAPI{
		FeedbacksAPI: feedbacks.NewFeedbacksAPI(auth),
	}
}
