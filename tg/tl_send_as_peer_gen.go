// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// SendAsPeer represents TL type `sendAsPeer#b81c7034`.
//
// See https://core.telegram.org/constructor/sendAsPeer for reference.
type SendAsPeer struct {
	// Flags field of SendAsPeer.
	Flags bin.Fields
	// PremiumRequired field of SendAsPeer.
	PremiumRequired bool
	// Peer field of SendAsPeer.
	Peer PeerClass
}

// SendAsPeerTypeID is TL type id of SendAsPeer.
const SendAsPeerTypeID = 0xb81c7034

// Ensuring interfaces in compile-time for SendAsPeer.
var (
	_ bin.Encoder     = &SendAsPeer{}
	_ bin.Decoder     = &SendAsPeer{}
	_ bin.BareEncoder = &SendAsPeer{}
	_ bin.BareDecoder = &SendAsPeer{}
)

func (s *SendAsPeer) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.PremiumRequired == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendAsPeer) String() string {
	if s == nil {
		return "SendAsPeer(nil)"
	}
	type Alias SendAsPeer
	return fmt.Sprintf("SendAsPeer%+v", Alias(*s))
}

// FillFrom fills SendAsPeer from given interface.
func (s *SendAsPeer) FillFrom(from interface {
	GetPremiumRequired() (value bool)
	GetPeer() (value PeerClass)
}) {
	s.PremiumRequired = from.GetPremiumRequired()
	s.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendAsPeer) TypeID() uint32 {
	return SendAsPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*SendAsPeer) TypeName() string {
	return "sendAsPeer"
}

// TypeInfo returns info about TL type.
func (s *SendAsPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendAsPeer",
		ID:   SendAsPeerTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PremiumRequired",
			SchemaName: "premium_required",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *SendAsPeer) SetFlags() {
	if !(s.PremiumRequired == false) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *SendAsPeer) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendAsPeer#b81c7034 as nil")
	}
	b.PutID(SendAsPeerTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendAsPeer) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendAsPeer#b81c7034 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendAsPeer#b81c7034: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode sendAsPeer#b81c7034: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendAsPeer#b81c7034: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SendAsPeer) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendAsPeer#b81c7034 to nil")
	}
	if err := b.ConsumeID(SendAsPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode sendAsPeer#b81c7034: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendAsPeer) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendAsPeer#b81c7034 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendAsPeer#b81c7034: field flags: %w", err)
		}
	}
	s.PremiumRequired = s.Flags.Has(0)
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendAsPeer#b81c7034: field peer: %w", err)
		}
		s.Peer = value
	}
	return nil
}

// SetPremiumRequired sets value of PremiumRequired conditional field.
func (s *SendAsPeer) SetPremiumRequired(value bool) {
	if value {
		s.Flags.Set(0)
		s.PremiumRequired = true
	} else {
		s.Flags.Unset(0)
		s.PremiumRequired = false
	}
}

// GetPremiumRequired returns value of PremiumRequired conditional field.
func (s *SendAsPeer) GetPremiumRequired() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (s *SendAsPeer) GetPeer() (value PeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}