package stratum

import (
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	"github.com/dubuqingfeng/stratum-server-monitor/stratum/btc"
	"github.com/dubuqingfeng/stratum-server-monitor/stratum/ckb"
	"github.com/dubuqingfeng/stratum-server-monitor/stratum/dcr"
)

func ParseHeight(coin string, resp interface{}) (height int64, err error) {
	nResp := resp.(models.NotifyRes)
	if coin == "dcr" {
		return dcr.ParseHeight(nResp.CoinbaseTX1)
	}
	if coin == "ckb" {
		return ckb.ParseHeight(nResp.Height)
	}
	if coin == "eth" || coin == "etc" {
		return ckb.ParseHeight(nResp.Height)
	}
	return btc.ParseHeight(nResp.CoinbaseTX1)
}

func ParsePrevHash(coin string, resp interface{}) (hash string) {
	nResp := resp.(models.NotifyRes)
	if coin == "ckb" {
		return ckb.ParsePrevHash(nResp.ParentHash)
	}
	return btc.ParsePrevHash(nResp.Hash)
}
