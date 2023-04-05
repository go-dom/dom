package lottery

import (
	"github.com/3JoB/ethclient"
	"github.com/3JoB/ethclient/rpc"
	errs "github.com/3JoB/ulib/err"
)

type Client struct {
	URL    string
	Option *rpc.ClientOption
	Client *ethclient.Client
	Debug  bool
}

// Initialize an client
func NewClient() *Client {
	return &Client{
		URL: "https://eth.rpc.rivet.cloud/",
	}
}

func (c *Client) SetUrl(url string) *Client {
	if url != "" {
		c.URL = url
	}
	return c
}

func (c *Client) SetDebug() *Client {
	if c.Debug {
		c.Debug = false
	} else {
		c.Debug = true
	}
	return c
}

func (c *Client) SetClientOption(option *rpc.ClientOption) *Client {
	if option != nil {
		c.Option = option
	}
	return c
}

// Create API connection
func (c *Client) Dial() (*Client, error) {
	var err error
	if c.Option != nil {
		c.Client, err = ethclient.Dial(c.URL, *c.Option)
	} else {
		c.Client, err = ethclient.Dial(c.URL)
	}
	return c, err
}

// Close the RPC client connection
func (c *Client) Close() {
	if c.Client != nil {
		c.Client.Close()
	}
}

var errDataEmpty error = &errs.Err{Op: "lottery.stream", Err: "data can not be empty!"}

// Create a lottery client
func (c *Client) NewStream(data *Data) (*Data, error) {
	if data == nil {
		return nil, errDataEmpty
	}
	data.client = c
	data.d = &d{}
	return data, nil
}

func (stream *Data) NewLotteryID() {
	stream.newLotteryID()
}