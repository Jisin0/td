// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// ChannelsGetFullChannelRequest represents TL type `channels.getFullChannel#8736a09`.
//
// See https://core.telegram.org/method/channels.getFullChannel for reference.
type ChannelsGetFullChannelRequest struct {
	// Channel field of ChannelsGetFullChannelRequest.
	Channel InputChannelClass
}

// ChannelsGetFullChannelRequestTypeID is TL type id of ChannelsGetFullChannelRequest.
const ChannelsGetFullChannelRequestTypeID = 0x8736a09

// Encode implements bin.Encoder.
func (g *ChannelsGetFullChannelRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getFullChannel#8736a09 as nil")
	}
	b.PutID(ChannelsGetFullChannelRequestTypeID)
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getFullChannel#8736a09: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getFullChannel#8736a09: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetFullChannelRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getFullChannel#8736a09 to nil")
	}
	if err := b.ConsumeID(ChannelsGetFullChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getFullChannel#8736a09: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getFullChannel#8736a09: field channel: %w", err)
		}
		g.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetFullChannelRequest.
var (
	_ bin.Encoder = &ChannelsGetFullChannelRequest{}
	_ bin.Decoder = &ChannelsGetFullChannelRequest{}
)

// ChannelsGetFullChannel invokes method channels.getFullChannel#8736a09 returning error if any.
//
// See https://core.telegram.org/method/channels.getFullChannel for reference.
func (c *Client) ChannelsGetFullChannel(ctx context.Context, request *ChannelsGetFullChannelRequest) (*MessagesChatFull, error) {
	var result MessagesChatFull
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}