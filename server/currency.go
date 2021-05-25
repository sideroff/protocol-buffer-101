package server

import (
	"context"
	"log"

	"github.com/sideroff/protocol-buffer-101/protocgen/currency"
)


type Currency struct {
	log *log.Logger
	currency.UnimplementedCurrencyServer // allows addition of methods without breaking builds because of unimplemented methods
}

func NewCurrency(l *log.Logger) *Currency {
	return &Currency{
		l,
		currency.UnimplementedCurrencyServer{},
	}
}

func (c *Currency) GetRate(ctx context.Context, req *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Printf("handle get rate, base: %s, destination: %s", req.GetBase(), req.GetDestination())


	res := &currency.RateResponse{
		Rate: 0.5,
	}

	return res, nil
}
