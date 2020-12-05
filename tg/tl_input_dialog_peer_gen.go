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

// InputDialogPeer represents TL type `inputDialogPeer#fcaafeb7`.
//
// See https://core.telegram.org/constructor/inputDialogPeer for reference.
type InputDialogPeer struct {
	// Peer field of InputDialogPeer.
	Peer InputPeerClass
}

// InputDialogPeerTypeID is TL type id of InputDialogPeer.
const InputDialogPeerTypeID = 0xfcaafeb7

// Encode implements bin.Encoder.
func (i *InputDialogPeer) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDialogPeer#fcaafeb7 as nil")
	}
	b.PutID(InputDialogPeerTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputDialogPeer#fcaafeb7: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputDialogPeer#fcaafeb7: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDialogPeer) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDialogPeer#fcaafeb7 to nil")
	}
	if err := b.ConsumeID(InputDialogPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDialogPeer#fcaafeb7: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputDialogPeer#fcaafeb7: field peer: %w", err)
		}
		i.Peer = value
	}
	return nil
}

// construct implements constructor of InputDialogPeerClass.
func (i InputDialogPeer) construct() InputDialogPeerClass { return &i }

// Ensuring interfaces in compile-time for InputDialogPeer.
var (
	_ bin.Encoder = &InputDialogPeer{}
	_ bin.Decoder = &InputDialogPeer{}

	_ InputDialogPeerClass = &InputDialogPeer{}
)

// InputDialogPeerFolder represents TL type `inputDialogPeerFolder#64600527`.
//
// See https://core.telegram.org/constructor/inputDialogPeerFolder for reference.
type InputDialogPeerFolder struct {
	// FolderID field of InputDialogPeerFolder.
	FolderID int
}

// InputDialogPeerFolderTypeID is TL type id of InputDialogPeerFolder.
const InputDialogPeerFolderTypeID = 0x64600527

// Encode implements bin.Encoder.
func (i *InputDialogPeerFolder) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDialogPeerFolder#64600527 as nil")
	}
	b.PutID(InputDialogPeerFolderTypeID)
	b.PutInt(i.FolderID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDialogPeerFolder) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDialogPeerFolder#64600527 to nil")
	}
	if err := b.ConsumeID(InputDialogPeerFolderTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDialogPeerFolder#64600527: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputDialogPeerFolder#64600527: field folder_id: %w", err)
		}
		i.FolderID = value
	}
	return nil
}

// construct implements constructor of InputDialogPeerClass.
func (i InputDialogPeerFolder) construct() InputDialogPeerClass { return &i }

// Ensuring interfaces in compile-time for InputDialogPeerFolder.
var (
	_ bin.Encoder = &InputDialogPeerFolder{}
	_ bin.Decoder = &InputDialogPeerFolder{}

	_ InputDialogPeerClass = &InputDialogPeerFolder{}
)

// InputDialogPeerClass represents InputDialogPeer generic type.
//
// See https://core.telegram.org/type/InputDialogPeer for reference.
//
// Example:
//  g, err := DecodeInputDialogPeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputDialogPeer: // inputDialogPeer#fcaafeb7
//  case *InputDialogPeerFolder: // inputDialogPeerFolder#64600527
//  default: panic(v)
//  }
type InputDialogPeerClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputDialogPeerClass
}

// DecodeInputDialogPeer implements binary de-serialization for InputDialogPeerClass.
func DecodeInputDialogPeer(buf *bin.Buffer) (InputDialogPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputDialogPeerTypeID:
		// Decoding inputDialogPeer#fcaafeb7.
		v := InputDialogPeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDialogPeerClass: %w", err)
		}
		return &v, nil
	case InputDialogPeerFolderTypeID:
		// Decoding inputDialogPeerFolder#64600527.
		v := InputDialogPeerFolder{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDialogPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputDialogPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputDialogPeer boxes the InputDialogPeerClass providing a helper.
type InputDialogPeerBox struct {
	InputDialogPeer InputDialogPeerClass
}

// Decode implements bin.Decoder for InputDialogPeerBox.
func (b *InputDialogPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputDialogPeerBox to nil")
	}
	v, err := DecodeInputDialogPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputDialogPeer = v
	return nil
}

// Encode implements bin.Encode for InputDialogPeerBox.
func (b *InputDialogPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputDialogPeer == nil {
		return fmt.Errorf("unable to encode InputDialogPeerClass as nil")
	}
	return b.InputDialogPeer.Encode(buf)
}