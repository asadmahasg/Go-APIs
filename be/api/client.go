package api

import (
	"project/libgo/common"
)

var c *Client

type Client struct {
	*common.CommonC
}

func Init() *Client {
	c = &Client{}
	c.CommonC = common.Init("api")

	return c
}
