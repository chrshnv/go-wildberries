package feedbacks

import (
	"time"
)

type WildberriesFeedbackListResponse struct {
	CountUnanswered int `json:"countUnanswered"`
	CountArchive    int `json:"countArchive"`
	Feedbacks       []struct {
		ID               string    `json:"ID"`
		Username         string    `json:"userName"`
		MatchingSize     string    `json:"matchingSize"`
		Text             string    `json:"text"`
		ProductValuation int       `json:"productValuation"`
		CreatedDate      time.Time `json:"createdDate"`
		State            string    `json:"state"`
		Answer           struct {
			Text  string `json:"text"`
			State string `json:"state"`
		} `json:"answer"`
		ProductDetails struct {
			SKU   string `json:"supplierArticle"`
			ImtId int    `json:"imtId"`
			NmId  int    `json:"nmId"`
		} `json:"productDetails"`
		PhotoLinks *[]struct {
			FullSize string `json:"fullSize"`
			MiniSize string `json:"miniSize"`
		} `json:"photoLinks"`
	} `json:"feedbacks"`
}

type WildberriesFeedbackPatchAnswer struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}
