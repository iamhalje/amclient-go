package amclient

import "fmt"

type Silence struct {
	ID        string    `json:"id,omitempty"`
	Matchers  []Matcher `json:"matchers"`
	StartsAt  string    `json:"startsAt"`
	EndsAt    string    `json:"endsAt"`
	CreatedBy string    `json:"createdBy"`
	Comment   string    `json:"comment"`
}

type Matcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}

func (c *Client) GetSilences() ([]Silence, error) {
	var silences []Silence
	err := c.Get("/api/v2/silences", &silences)
	return silences, err
}

func (c *Client) CreateSilence(silence Silence) (string, error) {
	var resp struct {
		SilenceID string `json:"silenceID"`
	}
	err := c.Post("/api/v2/silences", silence, &resp)
	return resp.SilenceID, err
}

func (c *Client) DeleteSilence(id string) error {
	return c.Delete(fmt.Sprintf("api/v2/silence/%s", id))
}
