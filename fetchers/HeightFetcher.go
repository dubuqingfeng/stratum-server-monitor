package fetchers

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dubuqingfeng/stratum-server-monitor/models"
	"github.com/dubuqingfeng/stratum-server-monitor/senders"
	"github.com/dubuqingfeng/stratum-server-monitor/stratum"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// by https://github.com/decred/gominer
// stratum server
type PoolHeightFetcher struct {
	Param    models.StratumServersParam
	Address  string
	Conn     net.Conn
	Reader   *bufio.Reader
	Height   int64
	PrevHash string
	ID       uint64
	AuthID   uint64
	SubID    uint64
	wg       sync.WaitGroup
}

var errJsonType = errors.New("unexpected type in json")

// start the monitor
func (p *PoolHeightFetcher) Start() {
	for {
		if err := p.Dial(); err != nil {
			p.HandleError(err)
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
	p.wg.Add(1)
	go p.Listen()
	p.Subscribe()
	p.Auth()
	p.wg.Wait()
}

// connect to stratum server
func (p *PoolHeightFetcher) Connect(limit int) error {
	err := p.Dial()
	if limit < 0 {
		return errors.New("limit")
	}
	if err != nil {
		p.HandleError(err)
		time.Sleep(5 * time.Second)
		return p.Connect(limit - 1)
	}
	return nil
}

func (p *PoolHeightFetcher) Reconnect() {
	p.ID = 1
	p.AuthID = 2
	err := p.Connect(3)
	if err != nil {
		p.HandleError(err)
		p.HandleError(errors.New("reconnect failed"))
		return
	}
	p.Subscribe()
	p.Auth()
}

// Dial
func (p *PoolHeightFetcher) Dial() error {
	var err error
	if p.Conn, err = net.Dial("tcp", p.Address); err != nil {
		return err
	}
	p.Reader = bufio.NewReader(p.Conn)
	return nil
}

// Subscribe to the event, https://gist.github.com/YihaoPeng/254d9daf3a5a80131507f32be6ed92df
func (p *PoolHeightFetcher) Subscribe() {
	msg := models.StratumMsg{Method: "mining.subscribe", ID: p.ID, Params: []string{"b-miner"}}
	p.SubID = msg.ID.(uint64)
	p.ID++
	p.WriteConn(msg)
}

// Auth by username and password
func (p *PoolHeightFetcher) Auth() {
	msg := models.StratumMsg{Method: "mining.authorize", ID: p.ID, Params: []string{p.Param.Username, p.Param.Password}}
	p.AuthID = msg.ID.(uint64)
	p.ID++
	p.WriteConn(msg)
}

// Write a message to the connection
func (p *PoolHeightFetcher) WriteConn(msg interface{}) {
	m, err := json.Marshal(msg)
	if err != nil {
		p.HandleError(err)
	}
	log.WithField("endpoint", p.Address).Info(string(m))
	if _, err := p.Conn.Write(m); err != nil {
		p.HandleError(err)
	}
	if _, err := p.Conn.Write([]byte("\n")); err != nil {
		p.HandleError(err)
	}
}

// Long connection event listening
func (p *PoolHeightFetcher) Listen() {
	defer p.wg.Done()
	log.Debug("Starting Listener")
	for {
		if p.Reader == nil {
			p.Reconnect()
		}
		result, err := p.Reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				p.HandleError(errors.New("get EOF. Connection lost! Reconnecting"))
				p.Reconnect()
			} else {
				p.HandleError(err)
			}
			time.Sleep(1 * time.Second)
			continue
		}
		log.WithField("endpoint", p.Address).Info(strings.TrimSuffix(result, "\n"))
		// Parse the returned data
		resp, err := p.Unmarshal([]byte(result))
		if err != nil {
			p.HandleError(err)
			continue
		}
		switch resp.(type) {
		case models.NotifyRes:
			p.handleNotifyRes(resp)
		case *models.SubscribeReply:
			p.handleSubscribeReply(resp)
		default:
			log.Debug("Unhandled message: ", result)
		}
	}
}

// Handle subscribe reply events
func (p *PoolHeightFetcher) handleSubscribeReply(resp interface{}) {
	log.Debug("Subscribe reply received.")
}

// Handle notify events
func (p *PoolHeightFetcher) handleNotifyRes(resp interface{}) {
	height, err := stratum.ParseHeight(p.Param.CoinType, resp)
	if err != nil {
		log.WithField("endpoint", p.Address).Errorf("failed to parse height %v", err)
	}
	prevHash := stratum.ParsePrevHash(p.Param.CoinType, resp)
	if height != p.Height {
		// The height has changed
		notification := &models.Notification{Height: height, OldHeight: p.Height, Reason: "", Username: p.Param.Username,
			Type: "HeightChanged", StratumServerURL: p.Address, CoinType: p.Param.CoinType,
			PrevHash: prevHash, StratumServerType: p.Param.Type, NotifiedAt: time.Now().UTC().String()}
		p.SendNotification(notification)
		log.WithField("endpoint", p.Address).Info(fmt.Sprintf("height: %d, old height: %d", height, p.Height))
	}

	// if height == p.Height && prevHash != p.PrevHash, create a notification.
	if height == p.Height && prevHash != p.PrevHash {
		notification := &models.Notification{Height: height, OldHeight: p.Height, Reason: "", Username: p.Param.Username,
			Type: "PrevHashChanged", StratumServerURL: p.Address, CoinType: p.Param.CoinType,
			PrevHash: prevHash, StratumServerType: p.Param.Type, NotifiedAt: time.Now().UTC().String()}
		p.SendNotification(notification)
		log.WithField("endpoint", p.Address).Info(fmt.Sprintf("height: %d, hash: %s, old hash: %s",
			height, p.PrevHash, prevHash))
	}

	// check coin base.
	if p.Param.CoinType == "ltc" || p.Param.CoinType == "btc" || p.Param.CoinType == "bch" {
		nResp := resp.(models.NotifyRes)
		blockStr := nResp.CoinbaseTX1 + "111111112222222222222222" + nResp.CoinbaseTX2
		if p.Param.Coinbase != "" {
			blockAddressMissing := strings.Index(blockStr, p.Param.Coinbase)
			if blockAddressMissing <= 0 {
				notification := &models.Notification{Height: height, OldHeight: p.Height, Reason: "", Username: p.Param.Username,
					Type: "CoinbaseAddressMissed", StratumServerURL: p.Address, CoinType: p.Param.CoinType,
					PrevHash: prevHash, StratumServerType: p.Param.Type, NotifiedAt: time.Now().UTC().String()}
				p.SendNotification(notification)
				log.WithField("endpoint", p.Address).Info(fmt.Sprintf("height: %d, old height: %d", height, p.Height))
			}
		}
		if p.Param.CoinbaseTags != "" {
			var CoinbaseTags map[string]interface{}
			err := json.Unmarshal([]byte(p.Param.CoinbaseTags), &CoinbaseTags)
			if err != nil {
				log.Error(err)
			} else {
				for key, value := range CoinbaseTags {
					blockAddressMissing := strings.Index(blockStr, value.(string))
					if blockAddressMissing <= 0 {
						notification := &models.Notification{Height: height, OldHeight: p.Height, Reason: "",
							Username: p.Param.Username, Type: "CoinbaseAddressMissed:" + key,
							StratumServerURL: p.Address, CoinType: p.Param.CoinType, PrevHash: prevHash,
							StratumServerType: p.Param.Type, NotifiedAt: time.Now().UTC().String()}
						p.SendNotification(notification)
						log.WithField("endpoint", p.Address).Info(fmt.Sprintf("height: %d, old height: %d", height, p.Height))
					}
				}
			}
		}
	}
	// mutex
	p.Height = height
	p.PrevHash = prevHash
}

// Unmarshal the message
func (p *PoolHeightFetcher) Unmarshal(blob []byte) (interface{}, error) {
	var (
		message map[string]json.RawMessage
		method  string
		id      uint64
		height  uint64
	)
	if err := json.Unmarshal(blob, &message); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(message["method"], &method); err != nil {
		method = ""
	}
	if err := json.Unmarshal(message["id"], &id); err != nil {
		return nil, err
	}
	if _, ok := message["height"]; ok {
		if err := json.Unmarshal(message["height"], &height); err != nil {
			return nil, err
		}
	}
	if id == p.AuthID {
		// {"id":2,"result":true,"error":null}
		// {"id":2,"result":null,"error":[29,"Invalid username",null]}
		var (
			result      bool
			errorHolder []interface{}
		)
		if err := json.Unmarshal(message["result"], &result); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(message["error"], &errorHolder); err != nil {
			return nil, err
		}
		resp := &models.BasicReply{ID: id, Result: result}
		if errorHolder != nil {
			var ok bool
			if resp.Error.ErrNum, ok = errorHolder[0].(uint64); !ok {
				return nil, errJsonType
			}
			if resp.Error.ErrStr, ok = errorHolder[1].(string); !ok {
				return nil, errJsonType
			}
		}
		return resp, nil
	}
	if id == p.SubID {
		// {"id":1,"result":[[["mining.set_difficulty","7fcc4632"],["mining.notify","7fcc4632"]],"7fcc4632",8],"error":null}
		var res []interface{}
		if err := json.Unmarshal(message["result"], &res); err != nil {
			return nil, err
		}
		if len(res) == 0 {
			return nil, errJsonType
		}
		resp := &models.SubscribeReply{}
		resp.ExtraNonce1 = res[1].(string)
		resp.ExtraNonce2Length = res[2].(float64)
		return resp, nil
	}

	switch method {
	case "mining.notify":
		var res []interface{}
		if err := json.Unmarshal(message["params"], &res); err != nil {
			return nil, err
		}
		nRes, err := p.BuildNotifyRes(res)
		if height != 0 {
			nRes.Height = float64(height)
		}
		return nRes, err

	case "mining.set_difficulty":
		// {"id":null,"method":"mining.set_difficulty","params":[64]}"
		var res []interface{}
		if err := json.Unmarshal(message["params"], &res); err != nil {
			return nil, err
		}
		difficulty, ok := res[0].(float64)
		if !ok {
			return nil, errJsonType
		}
		log.WithField("endpoint", p.Address).Infof("Stratum difficulty set to %v", difficulty)
		diffStr := strconv.FormatFloat(difficulty, 'E', -1, 32)
		var params []string
		params = append(params, diffStr)
		var nres = models.StratumMsg{Method: method, Params: params}
		return nres, nil

	default:
		resp := &models.StratumRsp{}
		err := json.Unmarshal(blob, &resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

func (p *PoolHeightFetcher) BuildNotifyRes(res []interface{}) (models.NotifyRes, error) {
	var nres = models.NotifyRes{}
	var ok bool
	if p.Param.CoinType == "ckb" {
		// https://wk588.com/forum.php?mod=viewthread&tid=19665
		// "jobId", "header hash", height, "parent hash", cleanJob
		if nres.JobID, ok = res[0].(string); !ok {
			return nres, errJsonType
		}
		if nres.Hash, ok = res[1].(string); !ok {
			return nres, errJsonType
		}
		if nres.Height, ok = res[2].(float64); !ok {
			return nres, errJsonType
		}
		if nres.ParentHash, ok = res[3].(string); !ok {
			return nres, errJsonType
		}
		if nres.CleanJobs, ok = res[4].(bool); !ok {
			return nres, errJsonType
		}
		return nres, nil
	}
	if p.Param.CoinType == "eth" || p.Param.CoinType == "etc" {
		if nres.Header, ok = res[0].(string); !ok {
			return nres, errJsonType
		}
		if nres.Header, ok = res[1].(string); !ok {
			return nres, errJsonType
		}
		if nres.Seed, ok = res[2].(string); !ok {
			return nres, errJsonType
		}
		if nres.ShareTarget, ok = res[3].(string); !ok {
			return nres, errJsonType
		}
		if nres.CleanJobs, ok = res[4].(bool); !ok {
			return nres, errJsonType
		}
		return nres, nil
	}
	// default: btc, ltc, dcr
	if nres.JobID, ok = res[0].(string); !ok {
		return nres, errJsonType
	}
	if nres.Hash, ok = res[1].(string); !ok {
		return nres, errJsonType
	}
	if nres.CoinbaseTX1, ok = res[2].(string); !ok {
		return nres, errJsonType
	}
	if nres.CoinbaseTX2, ok = res[3].(string); !ok {
		return nres, errJsonType
	}
	if nres.BlockVersion, ok = res[5].(string); !ok {
		return nres, errJsonType
	}
	if nres.Nbits, ok = res[6].(string); !ok {
		return nres, errJsonType
	}
	if nres.Ntime, ok = res[7].(string); !ok {
		return nres, errJsonType
	}
	if nres.CleanJobs, ok = res[8].(bool); !ok {
		return nres, errJsonType
	}
	return nres, nil
}

func (p *PoolHeightFetcher) SendNotifications(notifications []*models.Notification) {
	if len(notifications) == 0 {
		return
	}
	pushers := [2]senders.Sender{senders.SlackPusher, senders.MySQLPusher}
	for _, item := range pushers {
		if item == nil {
			continue
		}
		if !item.IsSupport() {
			continue
		}
		p.wg.Add(1)
		go func(sender senders.Sender, notifications []*models.Notification) {
			sender.Send(notifications)
			p.wg.Done()
		}(item, notifications)
	}
}

func (p *PoolHeightFetcher) SendNotification(notification *models.Notification) {
	// pusher list
	pushers := [2]senders.Sender{senders.SlackPusher, senders.MySQLPusher}
	for _, item := range pushers {
		if item == nil {
			continue
		}
		if !item.IsSupport() {
			continue
		}
		notifications := []*models.Notification{notification}
		go func(notifications []*models.Notification, sender senders.Sender) {
			sender.Send(notifications)
		}(notifications, item)
	}
}

func (p *PoolHeightFetcher) HandleError(err error) {
	log.WithField("endpoint", p.Address).Error(err)
}
