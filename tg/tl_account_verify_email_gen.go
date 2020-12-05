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

// AccountVerifyEmailRequest represents TL type `account.verifyEmail#ecba39db`.
//
// See https://core.telegram.org/method/account.verifyEmail for reference.
type AccountVerifyEmailRequest struct {
	// Email field of AccountVerifyEmailRequest.
	Email string
	// Code field of AccountVerifyEmailRequest.
	Code string
}

// AccountVerifyEmailRequestTypeID is TL type id of AccountVerifyEmailRequest.
const AccountVerifyEmailRequestTypeID = 0xecba39db

// Encode implements bin.Encoder.
func (v *AccountVerifyEmailRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode account.verifyEmail#ecba39db as nil")
	}
	b.PutID(AccountVerifyEmailRequestTypeID)
	b.PutString(v.Email)
	b.PutString(v.Code)
	return nil
}

// Decode implements bin.Decoder.
func (v *AccountVerifyEmailRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode account.verifyEmail#ecba39db to nil")
	}
	if err := b.ConsumeID(AccountVerifyEmailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.verifyEmail#ecba39db: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.verifyEmail#ecba39db: field email: %w", err)
		}
		v.Email = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.verifyEmail#ecba39db: field code: %w", err)
		}
		v.Code = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountVerifyEmailRequest.
var (
	_ bin.Encoder = &AccountVerifyEmailRequest{}
	_ bin.Decoder = &AccountVerifyEmailRequest{}
)

// AccountVerifyEmail invokes method account.verifyEmail#ecba39db returning error if any.
//
// See https://core.telegram.org/method/account.verifyEmail for reference.
func (c *Client) AccountVerifyEmail(ctx context.Context, request *AccountVerifyEmailRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}