package senders

import (
	"bytes"
	"fmt"
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// new bot : https://{team}.slack.com/services/new/bot
type SlackSender struct {
	Sender
}

type SlackMessage struct {
	ID      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	AsUser  bool   `json:"as_user"`
	Text    string `json:"text"`
	Token   string `json:"token"`
}

var SlackPusher SlackSender

func init() {
	SlackPusher = SlackSender{}
}

func (s SlackSender) IsSupport() bool {
	return utils.Config.SenderConfig.Slack.IsEnabled
}

func (s SlackSender) Send(notifications []*models.Notification) {
	if !utils.Config.SenderConfig.Slack.IsEnabled {
		return
	}
	for _, item := range notifications {
		if _, ok := utils.Config.BlackList[item.StratumServerURL]; ok {
			continue
		}
		if item.OldHeight >= item.Height {
			s.SingleSend(item)
		}
		if utils.Config.SenderConfig.Slack.SingleSendEnabled {
			s.SingleSend(item)
		}
	}
}

func (s SlackSender) SingleSend(notification *models.Notification) {
	s.SendText(s.BuildMessage(notification))
}

func (s SlackSender) SendText(text string) {
	message := SlackMessage{
		AsUser:  true,
		Channel: utils.Config.SenderConfig.Slack.Channel,
		Text:    text,
	}
	data := url.Values{}
	data.Set("token", utils.Config.SenderConfig.Slack.RobotToken)
	data.Add("channel", message.Channel)
	data.Add("text", message.Text)
	data.Add("as_user", strconv.FormatBool(message.AsUser))

	body, err := http.Post("https://slack.com/api/chat.postMessage", "application/x-www-form-urlencoded",
		bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Error(err)
		return
	}
	content, err := ioutil.ReadAll(body.Body)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(string(content))
}

func (s SlackSender) BuildMessage(notification *models.Notification) string {
	return fmt.Sprintf("height:%d,old height:%d,type:%s,monitor:%s,username:%s,ss:%s", notification.Height,
		notification.OldHeight, notification.Type, utils.Config.MonitorName, notification.Username,
		notification.StratumServerURL)
}
