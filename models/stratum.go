package models

import "encoding/json"

type StratumMsg struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      interface{} `json:"id"`
	JsonRPC string      `json:"jsonrpc,omitempty"`
	APIKey  string      `json:"api_key,omitempty"` // beam stratum 协议字段
}

type BasicReply struct {
	ID     interface{} `json:"id"`
	Error  StratumErr  `json:"error,omitempty"`
	Result bool        `json:"result"`
}

// StratumRsp is the basic response type from stratum.
type StratumRsp struct {
	Method string           `json:"method"`
	ID     interface{}      `json:"id"`
	Error  StratumErr       `json:"error,omitempty"`
	Result *json.RawMessage `json:"result,omitempty"`
	Height int64            `json:"height"` // beam
}

// NotifyRes models the json from a mining.notify message.
type NotifyRes struct {
	JobID          string
	Hash           string
	CoinbaseTX1    string
	CoinbaseTX2    string
	MerkleBranches []string
	BlockVersion   string
	Nbits          string
	Ntime          string
	CleanJobs      bool
	Height         float64 // ckb
	ParentHash     string  // ckb
	Seed           string  // eth | etc
	Header         string  // eth | etc
	ShareTarget    string  // eth | etc
}

type StratumErr struct {
	ErrNum float64
	ErrStr string
	Result *json.RawMessage `json:"result,omitempty"`
}

// SubscribeReply models the server response to a subscribe message.
type SubscribeReply struct {
	SubscribeID       string
	ExtraNonce1       string
	ExtraNonce2Length float64
}
