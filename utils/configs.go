package utils

import (
	"github.com/jinzhu/configor"
	"time"
)

var Config AppConfig

type AppConfig struct {
	Name      string `default:"app_name"`
	Debug     bool
	LogConfig struct {
		MaxAge time.Duration `default:"7"`
	}
	MonitorName                             string `default:"ss-monitor"`
	StratumServerConfigDatabaseTablePrefix  string
	StratumServerLogsDatabaseTablePrefix    string
	StratumServerMonitorDatabaseTablePrefix string
	GlobalDatabase                          MySQLDB
	StratumServerConfigDatabase             MySQLDB
	StratumServerLogsDatabase               MySQLDB
	StratumServerMonitorDatabase            MySQLDB
	SenderConfig                            SenderConfig
	SupportCoins                            []string
	CompareWithNodeEnabled                  bool
	BlackList                               map[string]string
}

// var Config = struct {
//
// }{}

type SenderConfig struct {
	BearyChat struct {
		IsEnabled      bool
		GroupEndpoint  string
		UnSupportTypes map[string]int
	}
	Alert struct {
		IsEnabled         bool
		SingleSendEnabled bool
		BaseURL           string
		AlertRoute        string
		ServiceName       string
		ServiceToken      string
		Channel           string
		UnSupportTypes    map[string]int
	}
	Slack struct {
		IsEnabled               bool
		SingleSendEnabled       bool
		RobotToken              string
		Channel                 string
		BlockTimeMonitorChannel string
		UnSupportTypes          map[string]int
	}
	MySQL struct {
		IsEnabled                        bool
		LogConfigDatabaseTablePrefix     string
		MonitorConfigDatabaseTablePrefix string
	}
}

type MySQLDSN struct {
	Name string
	DSN  string
}

type MySQLDB struct {
	Read     MySQLDSN
	Write    MySQLDSN
	Timezone string
}

func InitConfig(files string) {
	err := configor.Load(&Config, files)
	if err != nil {
		panic(err)
	}
}

func GetAllDatabaseConfigs() map[string]string {
	configs := make(map[string]string)
	AddDatabaseConfig(Config.GlobalDatabase, configs)
	AddDatabaseConfig(Config.StratumServerConfigDatabase, configs)
	AddDatabaseConfig(Config.StratumServerMonitorDatabase, configs)
	AddDatabaseConfig(Config.StratumServerLogsDatabase, configs)
	return configs
}

func AddDatabaseConfig(value MySQLDB, configs map[string]string) {
	if value.Read.DSN != "" && value.Read.Name != "" {
		configs[value.Read.Name] = value.Read.DSN
	}
	if value.Write.DSN != "" && value.Write.Name != "" {
		configs[value.Write.Name] = value.Write.DSN
	}
}
