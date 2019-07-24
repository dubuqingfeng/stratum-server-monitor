package models

import (
	"errors"
	"fmt"
	"github.com/dubuqingfeng/stratum-server-monitor/dbs"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	log "github.com/sirupsen/logrus"
)

type StratumServerHeight struct {
	Height           int64
	StratumServerURL string
	Type             string
	Username         string
	CoinType         string
	Description      string
	NotifiedAt       string
}

func GetStratumServerHeightsByCoinMySQL(coin string) ([]StratumServerHeight, error) {
	// select * from ss_heights where coin_type = "btc" order by height desc;
	conn := "ss:config:read"
	var list []StratumServerHeight
	if exists := dbs.CheckDBConnExists(conn); !exists {
		return list, errors.New("not found this database." + conn)
	}

	var sql string
	prefix := utils.Config.StratumServerConfigDatabaseTablePrefix
	sql = fmt.Sprintf("select height, stratum_server_url, type, username, coin_type, description, notified_at"+
		" from %s where coin_type = ? order by height desc;", prefix+"ss_heights")
	rows, err := dbs.DBMaps[conn].Query(sql, coin)
	if err != nil {
		log.Error(err)
		return list, err
	}
	for rows.Next() {
		var item StratumServerHeight
		if err := rows.Scan(&item.Height, &item.StratumServerURL, &item.Type, &item.Username, &item.CoinType,
			&item.Description, &item.NotifiedAt); err != nil {
			log.Error(err)
		}
		list = append(list, item)
	}

	if err := rows.Err(); err != nil {
		log.Error(err)
		return list, err
	}
	return list, nil
}
