package models

import (
	"errors"
	"fmt"
	"github.com/dubuqingfeng/stratum-server-monitor/dbs"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	log "github.com/sirupsen/logrus"
)

type StratumServersParam struct {
	ID           string
	Name         string
	Description  string
	Username     string
	Password     string
	CoinType     string
	Coinbase     string
	CoinbaseTags string
	Type         string
	Addresses    string
	Region       string
}

func GetAllStratumServersParams() ([]StratumServersParam, error) {
	conn := "ss:config:read"
	var list []StratumServersParam
	if exists := dbs.CheckDBConnExists(conn); !exists {
		return list, errors.New("not found this database." + conn)
	}

	var sql string
	prefix := utils.Config.StratumServerConfigDatabaseTablePrefix
	sql = fmt.Sprintf("select a.id, a.name, a.description, a.username, a.password, a.coin, a.coinbase, "+
		"a.coinbase_tags, a.type, b.addresses, b.region from %s a left join %s b on a.`server_id` = b.`id`;",
		prefix+"ss_params", prefix+"ss_servers")

	rows, err := dbs.DBMaps[conn].Query(sql)
	if err != nil {
		log.Error(err)
		return list, err
	}
	for rows.Next() {
		var param StratumServersParam
		if err := rows.Scan(&param.ID, &param.Name, &param.Description, &param.Username, &param.Password,
			&param.CoinType, &param.Coinbase, &param.CoinbaseTags, &param.Type, &param.Addresses, &param.Region); err != nil {
			log.Error(err)
		}
		list = append(list, param)
	}

	if err := rows.Err(); err != nil {
		log.Error(err)
		return list, err
	}
	return list, nil
}
