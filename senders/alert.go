package senders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// 自定义 Alert
type AlertSender struct {
	Sender
}

var AlertPusher AlertSender

func init() {
	AlertPusher = AlertSender{}
}

func (s AlertSender) IsSupport() bool {
	return utils.Config.SenderConfig.Alert.IsEnabled
}

func (s AlertSender) Send(notifications []*models.Notification) {
	if !utils.Config.SenderConfig.Alert.IsEnabled {
		return
	}
	for _, item := range notifications {
		if _, ok := utils.Config.BlackList[item.StratumServerURL]; ok {
			continue
		}
		if item.OldHeight >= item.Height {
			s.SingleSend(item)
		}
		if utils.Config.SenderConfig.Alert.SingleSendEnabled {
			s.SingleSend(item)
		}
	}
}

func (s AlertSender) SingleSend(notification *models.Notification) {
	s.SendText(s.BuildMessage(notification))
}

// send text
func (s AlertSender) SendText(text string) {
	message := map[string]interface{}{
		"action": "stratum_server_alert",
		"channel": utils.Config.SenderConfig.Alert.Channel,
		"title":  text,
		"message": text,
		"arg1": "",
		"arg2": "",
		"service_name": utils.Config.SenderConfig.Alert.ServiceName,
	}

	bytesMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{}
	request, err := http.NewRequest("POST", utils.Config.SenderConfig.Alert.BaseURL+
		utils.Config.SenderConfig.Alert.AlertRoute, bytes.NewBuffer(bytesMessage))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

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

func (s AlertSender) BuildMessage(notification *models.Notification) string {
	return fmt.Sprintf("height:%d,old height:%d,type:%s,monitor:%s,username:%s,ss:%s", notification.Height,
		notification.OldHeight, notification.Type, utils.Config.MonitorName, notification.Username,
		notification.StratumServerURL)
}
