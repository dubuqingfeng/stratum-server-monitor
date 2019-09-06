package cmd

import (
	"fmt"
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
		//root(cmd, args)
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

// initConfig reads in config file.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		utils.InitConfig(cfgFile)
	}
	dbs.InitMySQLDB()
}

func initLog() {
	log.SetLevel(log.InfoLevel)
}

