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

// HelpGetSupportRequest represents TL type `help.getSupport#9cdf08cd`.
//
// See https://core.telegram.org/method/help.getSupport for reference.
type HelpGetSupportRequest struct {
}

// HelpGetSupportRequestTypeID is TL type id of HelpGetSupportRequest.
const HelpGetSupportRequestTypeID = 0x9cdf08cd

// Encode implements bin.Encoder.
func (g *HelpGetSupportRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getSupport#9cdf08cd as nil")
	}
	b.PutID(HelpGetSupportRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetSupportRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getSupport#9cdf08cd to nil")
	}
	if err := b.ConsumeID(HelpGetSupportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getSupport#9cdf08cd: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetSupportRequest.
var (
	_ bin.Encoder = &HelpGetSupportRequest{}
	_ bin.Decoder = &HelpGetSupportRequest{}
)

// HelpGetSupport invokes method help.getSupport#9cdf08cd returning error if any.
//
// See https://core.telegram.org/method/help.getSupport for reference.
func (c *Client) HelpGetSupport(ctx context.Context, request *HelpGetSupportRequest) (*HelpSupport, error) {
	var result HelpSupport
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}