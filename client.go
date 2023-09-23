package dom

import (
	"github.com/3JoB/ethclient"
	"github.com/3JoB/ethclient/rpc"
)

type Client struct {
	URL    string            // ETH API Address
	Option *rpc.ClientOption // ETH client additional settings
	Client *ethclient.Client // ETH client
	Debug  bool              // Debug Mode
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

// Create a lottery Session
func (c *Client) NewSession(sessionData *Session) (*Session, error) {
	if sessionData == nil {
		return nil, ErrDataEmpty
	}
	sessionData.client = c
	sessionData.d = &d{}
	sessionData.UserNum = len(sessionData.UserID)
	return sessionData, nil
}

func (session *Session) NewLotteryID() {
	session.newLotteryID()
}
