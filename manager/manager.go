package manager

import (
	"encoding/json"
	"github.com/dubuqingfeng/stratum-server-monitor/fetchers"
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Manager struct {
	fetchers map[string][]*fetchers.PoolHeightFetcher
	// channel
	wg *sync.WaitGroup
}

func NewManager() *Manager {
	manager := &Manager{
		wg: &sync.WaitGroup{},
	}
	return manager
}

func (m Manager) Run() {
	// listen the add stratum server event.
	params, err := models.GetAllStratumServersParams()
	if err != nil {
		log.Error(err)
	}
	m.fetchers = make(map[string][]*fetchers.PoolHeightFetcher)
	for _, param := range params {
		var pools []string
		err := json.Unmarshal([]byte(param.Addresses), &pools)
		if err != nil {
			log.Error(err)
		}
		for _, pool := range pools {
			m.wg.Add(1)
			fetcher := fetchers.PoolHeightFetcher{Address: pool, Param: param, ID: 1, AuthID: 2}
			m.fetchers[param.CoinType] = append(m.fetchers[param.CoinType], &fetcher)
			go fetcher.Start()
		}
	}

	//go m.LoadAllStratumServers()
	m.wg.Add(1)
	go m.GetServersHeight()
	m.wg.Wait()
}

// channel receive
func (m Manager) AddServerParam() {

}

// Load the new stratum server configuration
func (m Manager) LoadAllStratumServers() {
}

// Gets the height of the stratum server
func (m Manager) GetServersHeight() {
	for {
		for coin, items := range m.fetchers {
			for _, item := range items {
				if item.Height > 0 {
					log.WithField("coin", coin).Debug(item.Height)
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}
