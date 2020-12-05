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

// WebPageAttributeTheme represents TL type `webPageAttributeTheme#54b56617`.
//
// See https://core.telegram.org/constructor/webPageAttributeTheme for reference.
type WebPageAttributeTheme struct {
	// Flags field of WebPageAttributeTheme.
	Flags bin.Fields
	// Documents field of WebPageAttributeTheme.
	//
	// Use SetDocuments and GetDocuments helpers.
	Documents []DocumentClass
	// Settings field of WebPageAttributeTheme.
	//
	// Use SetSettings and GetSettings helpers.
	Settings ThemeSettings
}

// WebPageAttributeThemeTypeID is TL type id of WebPageAttributeTheme.
const WebPageAttributeThemeTypeID = 0x54b56617

// Encode implements bin.Encoder.
func (w *WebPageAttributeTheme) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeTheme#54b56617 as nil")
	}
	b.PutID(WebPageAttributeThemeTypeID)
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field flags: %w", err)
	}
	if w.Flags.Has(0) {
		b.PutVectorHeader(len(w.Documents))
		for idx, v := range w.Documents {
			if v == nil {
				return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field documents element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field documents element with index %d: %w", idx, err)
			}
		}
	}
	if w.Flags.Has(1) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field settings: %w", err)
		}
	}
	return nil
}

// SetDocuments sets value of Documents conditional field.
func (w *WebPageAttributeTheme) SetDocuments(value []DocumentClass) {
	w.Flags.Set(0)
	w.Documents = value
}

// GetDocuments returns value of Documents conditional field and
// boolean which is true if field was set.
func (w *WebPageAttributeTheme) GetDocuments() (value []DocumentClass, ok bool) {
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.Documents, true
}

// SetSettings sets value of Settings conditional field.
func (w *WebPageAttributeTheme) SetSettings(value ThemeSettings) {
	w.Flags.Set(1)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WebPageAttributeTheme) GetSettings() (value ThemeSettings, ok bool) {
	if !w.Flags.Has(1) {
		return value, false
	}
	return w.Settings, true
}

// Decode implements bin.Decoder.
func (w *WebPageAttributeTheme) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeTheme#54b56617 to nil")
	}
	if err := b.ConsumeID(WebPageAttributeThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: %w", err)
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field flags: %w", err)
		}
	}
	if w.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field documents: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field documents: %w", err)
			}
			w.Documents = append(w.Documents, value)
		}
	}
	if w.Flags.Has(1) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for WebPageAttributeTheme.
var (
	_ bin.Encoder = &WebPageAttributeTheme{}
	_ bin.Decoder = &WebPageAttributeTheme{}
)