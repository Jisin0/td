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

// StatsAbsValueAndPrev represents TL type `statsAbsValueAndPrev#cb43acde`.
//
// See https://core.telegram.org/constructor/statsAbsValueAndPrev for reference.
type StatsAbsValueAndPrev struct {
	// Current field of StatsAbsValueAndPrev.
	Current float64
	// Previous field of StatsAbsValueAndPrev.
	Previous float64
}

// StatsAbsValueAndPrevTypeID is TL type id of StatsAbsValueAndPrev.
const StatsAbsValueAndPrevTypeID = 0xcb43acde

// Encode implements bin.Encoder.
func (s *StatsAbsValueAndPrev) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsAbsValueAndPrev#cb43acde as nil")
	}
	b.PutID(StatsAbsValueAndPrevTypeID)
	b.PutDouble(s.Current)
	b.PutDouble(s.Previous)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsAbsValueAndPrev) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsAbsValueAndPrev#cb43acde to nil")
	}
	if err := b.ConsumeID(StatsAbsValueAndPrevTypeID); err != nil {
		return fmt.Errorf("unable to decode statsAbsValueAndPrev#cb43acde: %w", err)
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode statsAbsValueAndPrev#cb43acde: field current: %w", err)
		}
		s.Current = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode statsAbsValueAndPrev#cb43acde: field previous: %w", err)
		}
		s.Previous = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsAbsValueAndPrev.
var (
	_ bin.Encoder = &StatsAbsValueAndPrev{}
	_ bin.Decoder = &StatsAbsValueAndPrev{}
)