package lottery

import (
	"github.com/3JoB/ethclient"
	"github.com/3JoB/ethclient/rpc"
	errs "github.com/3JoB/ulib/err"
)

type Client struct {
	URL    string // ETH API Address
	Option *rpc.ClientOption // ETH client additional settings
	Client *ethclient.Client // ETH client
	Debug  bool // Debug Mode
}

// Initialize an client
func NewClient() *Client {
	return &Client{
		URL: "https://eth.rpc.rivet.cloud/",
	}
}

// Set the ETH API address
func (c *Client) SetUrl(url string) *Client {
	if url != "" {
		c.URL = url
	}
	return c
}

// Set Debug mode
func (c *Client) SetDebug() *Client {
	if c.Debug {
		c.Debug = false
	} else {
		c.Debug = true
	}
	return c
}

// Set ETH client additional settings
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

var errDataEmpty error = &errs.Err{Op: "lottery.session", Err: "data can not be empty!"}

// Create a lottery Session
func (c *Client) NewSession(sessionData *Session) (*Session, error) {
	if sessionData == nil {
		return nil, errDataEmpty
	}
	sessionData.client = c
	sessionData.d = &d{}
	return sessionData, nil
}

func (session *Session) NewLotteryID() {
	session.newLotteryID()
}
