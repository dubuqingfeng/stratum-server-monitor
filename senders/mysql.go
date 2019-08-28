package senders

import (
	"errors"
	"fmt"
	"github.com/dubuqingfeng/stratum-server-monitor/dbs"
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type MySQLSender struct {
	Sender
}

var MySQLPusher MySQLSender

func init() {
	MySQLPusher = MySQLSender{}
}

func (s MySQLSender) IsSupport() bool {
	return utils.Config.SenderConfig.MySQL.IsEnabled
}

func (s MySQLSender) Send(notifications []*models.Notification) {
	if !utils.Config.SenderConfig.MySQL.IsEnabled {
		return
	}
	if len(notifications) <= 0 {
		return
	}
	// update server height
	if err := UpdateStratumServerHeight(notifications); err != nil {
		log.Error(err)
	}
	// batch insert
	if err := InsertNotificationsList(notifications); err != nil {
		log.Error(err)
	}
}

func InsertNotificationsList(list []*models.Notification) error {
	conn := utils.Config.StratumServerLogsDatabase.Write.Name
	exists := dbs.CheckDBConnExists(conn)
	if !exists {
		return errors.New("not found this database." + conn)
	}

	tableName := utils.Config.SenderConfig.MySQL.LogConfigDatabaseTablePrefix + "ss_logs"
	now := time.Now().UTC()
	stmtStrings := make([]string, 0, len(list))
	args := make([]interface{}, 0, len(list)*9)
	for _, item := range list {
		stmtStrings = append(stmtStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		args = append(args, item.Height)
		args = append(args, item.OldHeight)
		args = append(args, item.StratumServerURL)
		args = append(args, item.Type)
		args = append(args, item.Username)
		args = append(args, item.CoinType)
		args = append(args, item.NotifiedAt)
		args = append(args, now)
		args = append(args, now)
	}
	stmt := fmt.Sprintf("INSERT INTO `"+tableName+"` (height, old_height, stratum_server_url, type, username, " +
		"coin_type, notified_at, created_at, updated_at) VALUES %s", strings.Join(stmtStrings, ","))
	_, err := dbs.DBMaps[conn].Exec(stmt, args...)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// update stratum server height
func UpdateStratumServerHeight(list []*models.Notification) error {
	conn := utils.Config.StratumServerMonitorDatabase.Write.Name
	exists := dbs.CheckDBConnExists(conn)
	if !exists {
		return errors.New("not found this database." + conn)
	}
	tableName := utils.Config.SenderConfig.MySQL.MonitorConfigDatabaseTablePrefix + "ss_heights"
	for _, item := range list {
		var height int
		globalSql := "SELECT height FROM " + tableName + " WHERE stratum_server_url = ? and type = ? " +
			"and coin_type = ? ORDER BY height DESC LIMIT 0,1"
		err := dbs.DBMaps[conn].QueryRow(globalSql, item.StratumServerURL, item.StratumServerType, item.CoinType).Scan(&height)
		if err != nil {
			log.Info(err)
		}
		now := time.Now().UTC()
		if height > 0 && item.Height != 0 {
			// update
			stmt := fmt.Sprintf("UPDATE `" + tableName + "` set height = ?, notified_at = ?, updated_at = ? " +
				"where stratum_server_url = ? and type = ? and coin_type = ?")
			_, err = dbs.DBMaps[conn].Exec(stmt, item.Height, item.NotifiedAt, now, item.StratumServerURL,
				item.StratumServerType, item.CoinType)
			if err != nil {
				log.Error(err)
			}
			continue
		} else {
			stmt := fmt.Sprintf("INSERT INTO `" + tableName + "` (height, stratum_server_url, type, username, coin_type," +
				" notified_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
			_, err = dbs.DBMaps[conn].Exec(stmt, item.Height, item.StratumServerURL, item.StratumServerType, item.Username,
				item.CoinType, item.NotifiedAt, now, now)
			if err != nil {
				log.Error(err)
			}
		}
	}
	return nil
}
