package responses

import (
	"strings"
	"time"
)

type WildberriesResponse[T any] struct {
	Data             T        `json:"data"`
	Error            bool     `json:"error"`
	ErrorText        string   `json:"errorText"`
	AdditionalErrors []string `json:"additionalErrors"`
}

type CreatedDateTime time.Time

func (c *CreatedDateTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("yyyy-MM-dd'T'HH:mm:ssZ", value) //parse time
	if err != nil {
		return err
	}
	*c = CreatedDateTime(t) //set result using the pointer
	return nil
}

func (c CreatedDateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("yyyy-MM-dd'T'HH:mm:ssZ") + `"`), nil
}
