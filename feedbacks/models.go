package feedbacks

import (
	"github.com/chrshnv/go-wildberries/responses"
)

type WildberriesFeedbackListResponse struct {
	CountUnanswered int `json:"countUnanswered"`
	CountArchive    int `json:"countArchive"`
	Feedbacks       []struct {
		ID               string                    `json:"ID"`
		Username         string                    `json:"userName"`
		MatchingSize     string                    `json:"matchingSize"`
		Text             string                    `json:"text"`
		ProductValuation int                       `json:"productValuation"`
		CreatedDate      responses.CreatedDateTime `json:"createdDate"`
		State            string                    `json:"state"`
		Answer           struct {
			Text  string `json:"text"`
			State string `json:"state"`
		} `json:"answer"`
		ProductDetails struct {
			SKU string `json:"supplierArticle"`
		} `json:"productDetails"`
	} `json:"feedbacks"`
}
