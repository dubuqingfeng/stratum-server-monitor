package pool

import (
	"github.com/dubuqingfeng/stratum-server-monitor/fetchers"
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	"github.com/spf13/cobra"
)

var poolDemoArgs = struct {
	PoolAddress  string
	PoolName     string
	PoolUserName string
	PoolPassword string
	CoinType string
}{}

func NewRunPoolDemoCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "RunPoolDemo",
		Short: "",
		Run:   runPoolDemo,
	}
	cc.Flags().StringVar(&poolDemoArgs.PoolAddress, "address", "", "address")
	cc.Flags().StringVar(&poolDemoArgs.PoolName, "name", "", "name")
	cc.Flags().StringVar(&poolDemoArgs.PoolUserName, "username", "", "username")
	cc.Flags().StringVar(&poolDemoArgs.PoolPassword, "password", "", "password")
	cc.Flags().StringVar(&poolDemoArgs.CoinType, "coin", "btc", "coin")
	return cc
}

func runPoolDemo(cmd *cobra.Command, args []string) {
	param := models.StratumServersParam{Name: poolDemoArgs.PoolName, Username: poolDemoArgs.PoolUserName,
		Password: poolDemoArgs.PoolPassword, CoinType:poolDemoArgs.CoinType}
	fetcher := fetchers.PoolHeightFetcher{Address: poolDemoArgs.PoolAddress, Param: param, ID: 1, AuthID: 2}
	fetcher.Start()
}
