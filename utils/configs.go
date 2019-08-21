package utils

import "github.com/jinzhu/configor"

var Config = struct {
	Name                                   string `default:"app_name"`
	Debug                                  bool
	MonitorName                            string `default:"ss-monitor"`
	StratumServerConfigDatabaseTablePrefix string
	GlobalDatabase                         MySQLDB
	StratumServerConfigDatabase            MySQLDB
	SenderConfig                           SenderConfig
	SupportCoins                           []string
	BlackList                              map[string]string
}{}

type SenderConfig struct {
	BearyChat struct {
		IsEnabled      bool
		GroupEndpoint  string
		UnSupportTypes map[string]int
	}
	Slack struct {
		IsEnabled      bool
		RobotToken     string
		Channel        string
		UnSupportTypes map[string]int
	}
	MySQL struct {
		IsEnabled                    bool
		LogConfigDatabaseTablePrefix string
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
