package mockpackage

// External package usages are implemented using singleton client instances

// import (
// 	"url/mockpackage"
// )

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) DoSomething() error {
	return nil
}

