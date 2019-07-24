package main

import (
	"github.com/dubuqingfeng/stratum-server-monitor/cmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	cmd.Execute()
}
