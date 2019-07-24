package height

import (
	"github.com/dubuqingfeng/stratum-server-monitor/manager"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	"github.com/spf13/cobra"
	"time"
)

func NewHeightFetcherCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "HeightFetch",
		Short: "",
		Run:   stratumServerHeightFetch,
	}
	return cc
}

func stratumServerHeightFetch(cmd *cobra.Command, args []string) {
	utils.ConfigLocalFilesystemLogger("./logs/", "fetcher.log", 7*time.Hour*24, time.Second*20)
	// Initialization storage
	ssManager := manager.Manager{}
	ssManager.Run()
	quit := make(chan bool)
	<-quit
}
