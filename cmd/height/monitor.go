package height

import (
	"github.com/dubuqingfeng/stratum-server-monitor/monitors"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	"github.com/spf13/cobra"
	"time"
)

func NewHeightMonitorCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "HeightMonitor",
		Short: "",
		Run:   stratumServerHeightMonitor,
	}
	return cc
}

func stratumServerHeightMonitor(cmd *cobra.Command, args []string) {
	utils.ConfigLocalFilesystemLogger("./logs/", "monitor.log", 7*time.Hour*24, time.Second*20)
	// 查询
	monitor := monitors.HeightMonitor{}
	for {
		coins := monitor.GetSupportCoin()
		for _, coin := range coins {
			monitor.Run(coin)
		}
		time.Sleep(5 * time.Second)
	}
}
