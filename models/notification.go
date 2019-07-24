package models

type Notification struct {
	Height            int64
	OldHeight         int64
	Reason            string
	Type              string
	Username          string
	CoinType          string
	NotifiedAt        string
	StratumServerType string
	StratumServerURL  string
}
