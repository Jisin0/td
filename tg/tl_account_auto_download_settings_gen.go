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

// AccountAutoDownloadSettings represents TL type `account.autoDownloadSettings#63cacf26`.
//
// See https://core.telegram.org/constructor/account.autoDownloadSettings for reference.
type AccountAutoDownloadSettings struct {
	// Low field of AccountAutoDownloadSettings.
	Low AutoDownloadSettings
	// Medium field of AccountAutoDownloadSettings.
	Medium AutoDownloadSettings
	// High field of AccountAutoDownloadSettings.
	High AutoDownloadSettings
}

// AccountAutoDownloadSettingsTypeID is TL type id of AccountAutoDownloadSettings.
const AccountAutoDownloadSettingsTypeID = 0x63cacf26

// Encode implements bin.Encoder.
func (a *AccountAutoDownloadSettings) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode account.autoDownloadSettings#63cacf26 as nil")
	}
	b.PutID(AccountAutoDownloadSettingsTypeID)
	if err := a.Low.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.autoDownloadSettings#63cacf26: field low: %w", err)
	}
	if err := a.Medium.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.autoDownloadSettings#63cacf26: field medium: %w", err)
	}
	if err := a.High.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.autoDownloadSettings#63cacf26: field high: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AccountAutoDownloadSettings) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode account.autoDownloadSettings#63cacf26 to nil")
	}
	if err := b.ConsumeID(AccountAutoDownloadSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.autoDownloadSettings#63cacf26: %w", err)
	}
	{
		if err := a.Low.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.autoDownloadSettings#63cacf26: field low: %w", err)
		}
	}
	{
		if err := a.Medium.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.autoDownloadSettings#63cacf26: field medium: %w", err)
		}
	}
	{
		if err := a.High.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.autoDownloadSettings#63cacf26: field high: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountAutoDownloadSettings.
var (
	_ bin.Encoder = &AccountAutoDownloadSettings{}
	_ bin.Decoder = &AccountAutoDownloadSettings{}
)