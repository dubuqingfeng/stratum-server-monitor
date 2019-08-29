package stratum

import (
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	"github.com/dubuqingfeng/stratum-server-monitor/stratum/btc"
	"github.com/dubuqingfeng/stratum-server-monitor/stratum/dcr"
)

func ParseHeight(coin string, resp interface{}) (height int64, err error) {
	nResp := resp.(models.NotifyRes)
	if coin == "dcr" {
		return dcr.ParseHeight(nResp.CoinbaseTX1)
	}
	return btc.ParseHeight(nResp.CoinbaseTX1)
}

func ParsePrevHash(coin string, resp interface{}) (hash string) {
	nResp := resp.(models.NotifyRes)
	return btc.ParsePrevHash(nResp.Hash)
}