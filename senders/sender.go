package senders

import "github.com/dubuqingfeng/stratum-server-monitor/models"

type Sender interface {
	Send([]*models.Notification)
	IsSupport() bool
}
