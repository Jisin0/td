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

// MessagesAllStickersNotModified represents TL type `messages.allStickersNotModified#e86602c3`.
//
// See https://core.telegram.org/constructor/messages.allStickersNotModified for reference.
type MessagesAllStickersNotModified struct {
}

// MessagesAllStickersNotModifiedTypeID is TL type id of MessagesAllStickersNotModified.
const MessagesAllStickersNotModifiedTypeID = 0xe86602c3

// Encode implements bin.Encoder.
func (a *MessagesAllStickersNotModified) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.allStickersNotModified#e86602c3 as nil")
	}
	b.PutID(MessagesAllStickersNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAllStickersNotModified) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.allStickersNotModified#e86602c3 to nil")
	}
	if err := b.ConsumeID(MessagesAllStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.allStickersNotModified#e86602c3: %w", err)
	}
	return nil
}

// construct implements constructor of MessagesAllStickersClass.
func (a MessagesAllStickersNotModified) construct() MessagesAllStickersClass { return &a }

// Ensuring interfaces in compile-time for MessagesAllStickersNotModified.
var (
	_ bin.Encoder = &MessagesAllStickersNotModified{}
	_ bin.Decoder = &MessagesAllStickersNotModified{}

	_ MessagesAllStickersClass = &MessagesAllStickersNotModified{}
)

// MessagesAllStickers represents TL type `messages.allStickers#edfd405f`.
//
// See https://core.telegram.org/constructor/messages.allStickers for reference.
type MessagesAllStickers struct {
	// Hash field of MessagesAllStickers.
	Hash int
	// Sets field of MessagesAllStickers.
	Sets []StickerSet
}

// MessagesAllStickersTypeID is TL type id of MessagesAllStickers.
const MessagesAllStickersTypeID = 0xedfd405f

// Encode implements bin.Encoder.
func (a *MessagesAllStickers) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.allStickers#edfd405f as nil")
	}
	b.PutID(MessagesAllStickersTypeID)
	b.PutInt(a.Hash)
	b.PutVectorHeader(len(a.Sets))
	for idx, v := range a.Sets {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.allStickers#edfd405f: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAllStickers) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.allStickers#edfd405f to nil")
	}
	if err := b.ConsumeID(MessagesAllStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.allStickers#edfd405f: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.allStickers#edfd405f: field hash: %w", err)
		}
		a.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.allStickers#edfd405f: field sets: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerSet
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.allStickers#edfd405f: field sets: %w", err)
			}
			a.Sets = append(a.Sets, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesAllStickersClass.
func (a MessagesAllStickers) construct() MessagesAllStickersClass { return &a }

// Ensuring interfaces in compile-time for MessagesAllStickers.
var (
	_ bin.Encoder = &MessagesAllStickers{}
	_ bin.Decoder = &MessagesAllStickers{}

	_ MessagesAllStickersClass = &MessagesAllStickers{}
)

// MessagesAllStickersClass represents messages.AllStickers generic type.
//
// See https://core.telegram.org/type/messages.AllStickers for reference.
//
// Example:
//  g, err := DecodeMessagesAllStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesAllStickersNotModified: // messages.allStickersNotModified#e86602c3
//  case *MessagesAllStickers: // messages.allStickers#edfd405f
//  default: panic(v)
//  }
type MessagesAllStickersClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesAllStickersClass
}

// DecodeMessagesAllStickers implements binary de-serialization for MessagesAllStickersClass.
func DecodeMessagesAllStickers(buf *bin.Buffer) (MessagesAllStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesAllStickersNotModifiedTypeID:
		// Decoding messages.allStickersNotModified#e86602c3.
		v := MessagesAllStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesAllStickersClass: %w", err)
		}
		return &v, nil
	case MessagesAllStickersTypeID:
		// Decoding messages.allStickers#edfd405f.
		v := MessagesAllStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesAllStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesAllStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesAllStickers boxes the MessagesAllStickersClass providing a helper.
type MessagesAllStickersBox struct {
	AllStickers MessagesAllStickersClass
}

// Decode implements bin.Decoder for MessagesAllStickersBox.
func (b *MessagesAllStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesAllStickersBox to nil")
	}
	v, err := DecodeMessagesAllStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AllStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesAllStickersBox.
func (b *MessagesAllStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AllStickers == nil {
		return fmt.Errorf("unable to encode MessagesAllStickersClass as nil")
	}
	return b.AllStickers.Encode(buf)
}