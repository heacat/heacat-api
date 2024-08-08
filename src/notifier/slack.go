package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/heacat/heacat-api/src/config"
	"github.com/heacat/heacat-api/src/logger"
)

type webhook_post_t struct {
	Text string `json:"text"`
}

func SlackNotifier(message string) {
	data := webhook_post_t{
		Text: message,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Log.Errorf("Error while marshalling data: %s", err)
		return
	}

	resp, err := http.Post(config.Config.Alarm.Slack.WebHookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Log.Errorf("Error while sending message to slack: %s", err)
		return
	}
	defer resp.Body.Close()
}
