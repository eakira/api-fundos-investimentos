package externo

import (
	"github.com/go-resty/resty/v2"
)

var (
	client  *resty.Client
	CVM_URL = "CVM_URL"
)

type fundosClient struct{}

func NewFundosClient() *fundosClient {
	return &fundosClient{}
}
