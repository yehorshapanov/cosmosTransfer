package cosmos_api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type MsgValue struct {
	FromAddress string      `json:"from_address"`
	ToAddress   string      `json:"to_address"`
	Amount      interface{} `json:"amount"`
}

type Msg struct {
	Type  string   `json:"type"`
	Value MsgValue `json:"value"`
}

type AmountObject struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Fee struct {
	Amount []AmountObject `json:"amount"`
	Gas    string         `json:"gas"`
}

type Value struct {
	Msg []Msg `json:"msg"`
	Fee Fee   `json:"fee"`
}

type TxDetails struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}

type Event struct {
	Type string `json:"type"`
}

type Tx struct {
	Hash      string    `json:"txhash"`
	TxDetails TxDetails `json:"tx"`
	Events    []Event   `json:"events"`
	Timestamp time.Time `json:"timestamp"`
}

func (t Tx) GetAmount() string {
	if t.TxDetails.Value.Msg == nil {
		return "0"
	}

	if amount, ok := t.TxDetails.Value.Msg[0].Value.Amount.([]interface{}); ok {
		tmp := amount[0].(map[string]interface{})
		return fmt.Sprint(tmp["amount"])
	}

	if amount, ok := t.TxDetails.Value.Msg[0].Value.Amount.(map[string]interface{}); ok {
		return fmt.Sprint(amount["amount"])
	}

	return ""
}

func (t Tx) GetFromAddress() string {
	if t.TxDetails.Value.Msg == nil {
		return "0"
	}

	return t.TxDetails.Value.Msg[0].Value.FromAddress
}

func (t Tx) GetToAddress() string {
	if t.TxDetails.Value.Msg == nil {
		return "0"
	}

	return t.TxDetails.Value.Msg[0].Value.ToAddress
}

func (t Tx) GetFee() string {
	if t.TxDetails.Value.Fee.Amount == nil {
		return "0"
	}

	return t.TxDetails.Value.Fee.Amount[0].Amount
}

func (c *Client) GetTx(ctx context.Context, txHash string) (*Tx, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/txs/%s", c.BaseURL, txHash), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Tx{}
	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
