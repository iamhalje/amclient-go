package amclient

import (
	"fmt"
	"time"
)

type Silence struct {
	ID        string    `json:"id,omitempty"`
	Status    Status    `json:"status"`
	Matchers  []Matcher `json:"matchers"`
	StartsAt  string    `json:"startsAt"`
	EndsAt    string    `json:"endsAt"`
	CreatedBy string    `json:"createdBy"`
	Comment   string    `json:"comment"`
}

type Status struct {
	State string `json:"state"`
}

type Matcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
	IsEqual bool   `json:"isEqual"`
}

func (c *Client) GetSilences() ([]Silence, error) {
	var silences []Silence
	err := c.Get("/api/v2/silences", &silences)

	if err != nil {
		return nil, err
	}

	var activeSilences []Silence
	now := time.Now()

	// Get only Active and Pending
	for _, s := range silences {
		if s.Status.State == "expired" {
			continue
		}

		endsAtTime, err := time.Parse(time.RFC3339, s.EndsAt)
		if err != nil {
			fmt.Printf("failed to parse EndsAt for silence %s: %v\n", s.ID, err)
			continue
		}

		if endsAtTime.Before(now) {
			continue
		}

		activeSilences = append(activeSilences, s)
	}

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
	return c.Delete(fmt.Sprintf("/api/v2/silence/%s", id))
}
