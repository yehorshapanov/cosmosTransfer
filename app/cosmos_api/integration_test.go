//go:build integration
// +build integration

package cosmos_api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBlock(t *testing.T) {
	c := NewClient()

	ctx := context.Background()
	res, err := c.GetBlock(ctx, 9989379)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.NotEmpty(t, res.BlockId.Hash, "expecting non-empty hash")
}

func TestGetTx(t *testing.T) {
	c := NewClient()

	res, err := c.GetTx(context.Background(), "8B670BA2F1A98AE133532C4856AC87F0185ADF3980058688803BC70682897448")

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.NotEmpty(t, res.TxDetails.Type)
}
