package height

import (
	"github.com/dubuqingfeng/stratum-server-monitor/manager"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	_ "net/http/pprof"
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
	utils.ConfigLocalFilesystemLogger("./logs/", "fetcher.log",
		utils.Config.LogConfig.MaxAge*time.Hour*24, time.Second*20)
	if utils.Config.Debug {
		go func() {
			if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
				log.Error(err)
			}
		}()
	}
	ssManager := manager.NewManager()
	ssManager.Run()
	quit := make(chan bool)
	<-quit
}
