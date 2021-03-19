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

}

func (c *Client) GetExchange() (Exchange, error) {

}
