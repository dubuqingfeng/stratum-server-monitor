package dbs

import (
	"database/sql"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	_ "github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
	"time"
)

var (
	DBMaps map[string]*sql.DB
)

// Initialize database
func InitMySQLDB() {
	// Initialize all mysql connections
	DBMaps = make(map[string]*sql.DB)
	configs := utils.GetAllDatabaseConfigs()
	for k, v := range configs {
		tempDB, err := sql.Open("mysql", v)
		if err != nil {
			log.Error(err)
			err := tempDB.Close()
			if err != nil {
				log.Error(err)
				panic(err.Error())
			}
			panic(err.Error())
		}
		tempDB.SetConnMaxLifetime(time.Minute * 10)
		tempDB.SetMaxIdleConns(10)
		tempDB.SetMaxOpenConns(20)
		DBMaps[k] = tempDB
	}
}

// Check if the connection exists
func CheckDBConnExists(conn string) bool {
	if _, ok := DBMaps[conn]; ok {
		return true
	}
	return false
}
