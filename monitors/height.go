package monitors

import (
	"fmt"
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	"github.com/dubuqingfeng/stratum-server-monitor/senders"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type HeightMonitor struct {
}

func (h *HeightMonitor) GetSupportCoin() []string {
	return utils.Config.SupportCoins
}

func (h *HeightMonitor) Run(coin string) {
	var text string
	var height int64
	var latest models.StratumServerHeight
	list, err := models.GetStratumServerHeightsByCoinMySQL(coin)
	if err != nil {
		log.Error(err)
	}

	for _, item := range list {
		if _, ok := utils.Config.BlackList[item.StratumServerURL]; ok {
			continue
		}
		if height == 0 {
			height = item.Height
			latest = item
		}
		ssURLPort := strings.Split(item.StratumServerURL, ":")
		if len(ssURLPort) > 1 && ssURLPort[1] == "25" {
			continue
		}
		if item.Height < height {
			text += fmt.Sprintf("ss_url: %s ，高度：%d，高度差：%d\n", item.StratumServerURL, item.Height,
				height-item.Height)
		}
	}
	// Compared with nodes do monitoring
	if utils.Config.CompareWithNodeEnabled {
		node, err := models.GetPeerHeightsByCoinMySQL(coin)
		if err != nil {
			log.Error(err)
		}
		if node.Height+1 > latest.Height {
			text += fmt.Sprintf("node: %s，高度：%d，高度差：%d\n", node.StratumServerURL,
				node.Height, node.Height+1-latest.Height)
		}
	}
	if text == "" {
		return
	}
	text += fmt.Sprintf("最新高度：%d，ss_url: %s \n", height, latest.StratumServerURL)
	text += fmt.Sprintf("抓取时间：%s，Monitor：%s", time.Now().Format("2006-01-02 15:04:05"),
		utils.Config.MonitorName)
	log.Info(text)
	channel := utils.Config.SenderConfig.Slack.BlockTimeMonitorChannel
	senders.SlackPusher.SendText(channel, "", text)
}
