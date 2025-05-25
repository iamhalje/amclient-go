# amclient-go

Go client for the [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/).

## Example

```go
client := amclient.NewClient("http://localhost:9093")

// Get alerts
alerts, _ := client.GetAlerts()

// Get status
status, _ := client.GetStatus()

// Get silences
silences, _ := client.GetSilences()

// Create silence
silence := amclient.Silence{
    Matchers: []amclient.Matcher{
        {Name: "alertname", Value: "TestAlert"},
    },
    StartsAt:  "2025-05-25T00:00:00Z",
    EndsAt:    "2025-05-26T00:00:00Z",
    CreatedBy: "user",
    Comment:   "test silence",
}
id, _ := client.CreateSilence(silence)

// Delete silence
_ = client.DeleteSilence(id)
```
