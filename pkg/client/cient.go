package client

import "github.com/sarthakjdev/wapi.go/pkg/manager"

type Client struct {
	media     manager.MediaManager
	message   manager.MessageManager
	requester requestClient
}

func NewClient() *Client {
	return &Client{
		media:     *manager.NewMediaManager(),
		message:   *manager.NewMessageManager(),
		requester: *NewRequestClient(),
	}
}

