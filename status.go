package amclient

type StatusResponse struct {
	VersionInfo struct {
		Version   string `json:"version"`
		BuildDate string `json:"buildDate"`
		GoVersion string `json:"goVersion"`
	} `json:"versionInfo"`
	Uptime string `json:"update"`
}

func (c *Client) GetStatus() (*StatusResponse, error) {
	var status StatusResponse
	err := c.Get("/api/v2/status", &status)
	return &status, err
}
