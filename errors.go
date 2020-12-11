package bittrex

import "errors"

var(
	ERR_ORDER_MISSING_PARAMETERS = errors.New("missing parameters. make sure (type, market_symbol, direction, time_in_force) are set")
	ERR_WITHDRAWAL_MISSING_PARAMETERS = errors.New("missing parameters. make sure (address, currency, quantity) are set")
)

