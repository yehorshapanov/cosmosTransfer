package cosmos_api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Parts struct {
	Total int    `json:"total"`
	Hash  string `json:"hash"`
}

type BlockId struct {
	Hash  string `json:"hash"`
	Parts Parts  `json:"parts"`
}

type Data struct {
	Txs []string `json:"txs"`
}

type Header struct {
	Height          string    `json:"height"`
	Time            time.Time `json:"time"`
	ProposerAddress string    `json:"proposer_address"`
}

type Block struct {
	Data   Data   `json:"data"`
	Header Header `json:"header"`
}

type BlockObject struct {
	BlockId BlockId `json:"block_id"`
	Block   Block   `json:"block"`
}

func (c *Client) GetBlock(ctx context.Context, blockNumber int) (*BlockObject, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocks/%d", c.BaseURL, blockNumber), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := BlockObject{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
