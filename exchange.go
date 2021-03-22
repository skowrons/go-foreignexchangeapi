package fera

import "time"

// Exchange holds data about the current exchanges.
// The standard Base is EUR but can be overriden by the user.
type Exchange struct {
	Base string
	Date time.Time
	Rates []Rate
}

type Rate struct {
	CurrencyShort string
	Value float32
}

type ExchangeOptions struct {
	Base string
	Symbols []Symbol
	StartAt time.Time
	EndAt time.Time
}

// GetExchange will return the current exchange, when no starting and endpoint is given.
// If a starting and endpoint is given the function will return histroy data.
func (c *Client) GetExchange(opt ExchangeOptions) (*Exchange, error) {
	return nil, nil
}
