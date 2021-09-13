package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/dubuqingfeng/stratum-server-monitor/cmd/height"
	"github.com/dubuqingfeng/stratum-server-monitor/cmd/pool"
	"github.com/dubuqingfeng/stratum-server-monitor/dbs"
	"github.com/dubuqingfeng/stratum-server-monitor/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ss-monitor",
	Short: "",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// root(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initApplication)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./configs/config.yaml", "config file (default is $HOME/.ethAdmin.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(
		height.NewHeightFetcherCommand(),
		height.NewHeightMonitorCommand(),
		pool.NewRunPoolDemoCommand(),
	)
}

func initApplication() {
	initConfig()
	initLog()
}

type ChangeListener struct {
}

func (c *ChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	fmt.Println(event.Changes)
}

func (c *ChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	fmt.Println("change listener.")
	fmt.Println(changeEvent.Changes)
	fmt.Println(changeEvent.Namespace)
}

func initApolloConfig() {
	namespaceName := "stratum-server-montior.json"
	readyConfig := &config.AppConfig{
		IsBackupConfig:   true,
		BackupConfigPath: "./",
		AppID:            utils.GetEnv("APOLLO_CUSTOM_CONFIG_APP_ID", ""),
		Cluster:          utils.GetEnv("APOLLO_CUSTOM_CONFIG_CLUSTER_NAME", "default"),
		NamespaceName:    namespaceName,
		IP:               utils.GetEnv("APOLLO_CUSTOM_CONFIG_SERVICE_IP", ""),
	}
	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return readyConfig, nil
	})
	c := client.GetConfig(namespaceName)
	if err := json.Unmarshal([]byte(c.GetValue("content")), &utils.Config); err != nil {
		log.Error(err)
	}
	log.WithField("isInit", c.GetIsInit()).Info("apollo init completed.")
	client.AddChangeListener(&ChangeListener{})
}

// initConfig reads in config file.
func initConfig() {
	configDriverName := os.Getenv("CONFIG_DRIVER_NAME")
	if configDriverName == "apollo" {
		initApolloConfig()
	} else {
		if cfgFile != "" {
			// Use config file from the flag.
			utils.InitConfig(cfgFile)
		}
	}
	dbs.InitMySQLDB()
}

func initLog() {
	log.SetLevel(log.InfoLevel)
}
