package balance

import (
	"net/http"

	messagebird "github.com/Bijles-aan-Huis-B-V/messagebird-go-rest-api"
)

// Balance describes your balance information.
type Balance struct {
	Payment string
	Type    string
	Amount  float32
}

const path = "balance"

// Read returns the balance information for the account that is associated with
// the access key.
func Read(c messagebird.Client) (*Balance, error) {
	balance := &Balance{}
	if err := c.Request(balance, http.MethodGet, path, nil); err != nil {
		return nil, err
	}

	return balance, nil
}
