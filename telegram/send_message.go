package telegram

import (
	"context"

	"github.com/gotd/td/tg"
)

// SendMessage sends message to peer.
func (c *Client) SendMessage(ctx context.Context, req *tg.MessagesSendMessageRequest) error {
	updates, err := c.tg.MessagesSendMessage(ctx, req)
	if err != nil {
		return err
	}
	return c.processUpdates(updates)
}