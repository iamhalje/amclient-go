package amclient

type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    string            `json:"startsAt"`
	EndsAt      string            `json:"endsAt"`
	Status      struct {
		State string `json:"state"`
	} `json:"status"`
}

func (c *Client) GetAlerts() ([]Alert, error) {
	var alerts []Alert
	err := c.Get("/api/v2/alerts", &alerts)
	return alerts, err
}
