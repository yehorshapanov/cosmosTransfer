package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const tx = "ClYKUgobL2Nvc21vcy5nb3YudjFiZXRhMS5Nc2dWb3RlEjMIQRItY29zbW9zMXljNmczcHF1YzVucTd6cWRsNnVsNGR4bnhmaDc3NHd1dWVxMmRnGAESABJnClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiED3QPf5hB9k70tomhgqxWAD53Hq/Kidgc0ZMjywHlEgQUSBAoCCH8YFhITCg0KBXVhdG9tEgQ2MjUwEJChDxpA5HqQHjv+592qZET4/IznVRH8GH2rFXQpA6eAEIlgr1Q9sV8XGnsmLqRv7nPAqvyyt5sEC9nA6FG047GacIXQiA=="
const exp = "8B670BA2F1A98AE133532C4856AC87F0185ADF3980058688803BC70682897448"

func TestCalculateTransactionId(t *testing.T) {
	res := CalculateTransactionID(tx)

	assert.Equal(t, res, exp)
}
