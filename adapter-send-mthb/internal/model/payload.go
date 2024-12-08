package model

import "math/big"

type Payload struct {
	Amount           big.Int `json:"amount"`
	To               string  `json:"to"`
	From             string  `json:"from"`
	IssuerPrivateKey string  `json:"issuerPrivateKey"`
}

// {"amount": 1, "to": "0x85b9D2c19B9fb1eA646D0D0EeCCA4d9e78AB6E74","from" : "0x6D928aEa1140E7A8fAa178f6032d0950cc1D1F80", "issuerPrivateKey" : "9668ffc3d3337c62146bea48a88a61c7b6f7a35844aa30ca1789d4bd40496855"}
