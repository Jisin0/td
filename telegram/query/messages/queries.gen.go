// Code generated by itergen, DO NOT EDIT.

package messages

import (
	"context"

	"github.com/ogen-go/errors"

	"github.com/gotd/td/tg"
)

// No-op definition for keeping imports.
var _ = context.Background()

// Request is a parameter for Query.
type Request struct {
	AddOffset  int
	OffsetDate int
	OffsetID   int
	OffsetPeer tg.InputPeerClass
	OffsetRate int
	Limit      int
}

// Query is an abstraction for messages request.
// NB: iterator mutates returned data (sorts, at least).
type Query interface {
	Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error)
}

// QueryFunc is a function adapter for Query.
type QueryFunc func(ctx context.Context, req Request) (tg.MessagesMessagesClass, error)

// Query implements Query interface.
func (q QueryFunc) Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error) {
	return q(ctx, req)
}

// QueryBuilder is a helper to create message queries.
type QueryBuilder struct {
	raw *tg.Client
}

// NewQueryBuilder creates new QueryBuilder.
func NewQueryBuilder(raw *tg.Client) *QueryBuilder {
	return &QueryBuilder{raw: raw}
}

// GetHistoryQueryBuilder is query builder of MessagesGetHistory.
type GetHistoryQueryBuilder struct {
	raw        *tg.Client
	req        tg.MessagesGetHistoryRequest
	batchSize  int
	addOffset  int
	offsetDate int
	offsetID   int
}

// GetHistory creates query builder of MessagesGetHistory.
func (q *QueryBuilder) GetHistory(paramPeer tg.InputPeerClass) *GetHistoryQueryBuilder {
	b := &GetHistoryQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req: tg.MessagesGetHistoryRequest{
			Peer: &tg.InputPeerEmpty{},
		},
	}

	b.req.Peer = paramPeer
	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *GetHistoryQueryBuilder) BatchSize(batchSize int) *GetHistoryQueryBuilder {
	b.batchSize = batchSize
	return b
}

// OffsetDate sets offsetDate from which iterate start.
func (b *GetHistoryQueryBuilder) OffsetDate(offsetDate int) *GetHistoryQueryBuilder {
	b.offsetDate = offsetDate
	return b
}

// OffsetID sets offsetID from which iterate start.
func (b *GetHistoryQueryBuilder) OffsetID(offsetID int) *GetHistoryQueryBuilder {
	b.offsetID = offsetID
	return b
}

// Peer sets Peer field of GetHistory query.
func (b *GetHistoryQueryBuilder) Peer(paramPeer tg.InputPeerClass) *GetHistoryQueryBuilder {
	b.req.Peer = paramPeer
	return b
}

// Query implements Query interface.
func (b *GetHistoryQueryBuilder) Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error) {
	r := &tg.MessagesGetHistoryRequest{
		Limit: req.Limit,
	}

	r.Peer = b.req.Peer
	r.AddOffset = req.AddOffset
	r.OffsetDate = req.OffsetDate
	r.OffsetID = req.OffsetID
	return b.raw.MessagesGetHistory(ctx, r)
}

// Iter returns iterator using built query.
func (b *GetHistoryQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	iter = iter.OffsetDate(b.offsetDate)
	iter = iter.OffsetID(b.offsetID)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *GetHistoryQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}

// Count fetches remote state to get number of elements.
func (b *GetHistoryQueryBuilder) Count(ctx context.Context) (int, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get total")
	}
	return c, nil
}

// Collect creates iterator and collects all elements to slice.
func (b *GetHistoryQueryBuilder) Collect(ctx context.Context) ([]Elem, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	r := make([]Elem, 0, c)
	for iter.Next(ctx) {
		r = append(r, iter.Value())
	}

	return r, iter.Err()
}

// GetRecentLocationsQueryBuilder is query builder of MessagesGetRecentLocations.
type GetRecentLocationsQueryBuilder struct {
	raw       *tg.Client
	req       tg.MessagesGetRecentLocationsRequest
	batchSize int
}

// GetRecentLocations creates query builder of MessagesGetRecentLocations.
func (q *QueryBuilder) GetRecentLocations(paramPeer tg.InputPeerClass) *GetRecentLocationsQueryBuilder {
	b := &GetRecentLocationsQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req: tg.MessagesGetRecentLocationsRequest{
			Peer: &tg.InputPeerEmpty{},
		},
	}

	b.req.Peer = paramPeer
	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *GetRecentLocationsQueryBuilder) BatchSize(batchSize int) *GetRecentLocationsQueryBuilder {
	b.batchSize = batchSize
	return b
}

// Peer sets Peer field of GetRecentLocations query.
func (b *GetRecentLocationsQueryBuilder) Peer(paramPeer tg.InputPeerClass) *GetRecentLocationsQueryBuilder {
	b.req.Peer = paramPeer
	return b
}

// Query implements Query interface.
func (b *GetRecentLocationsQueryBuilder) Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error) {
	r := &tg.MessagesGetRecentLocationsRequest{
		Limit: req.Limit,
	}

	r.Peer = b.req.Peer
	return b.raw.MessagesGetRecentLocations(ctx, r)
}

// Iter returns iterator using built query.
func (b *GetRecentLocationsQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *GetRecentLocationsQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}

// Count fetches remote state to get number of elements.
func (b *GetRecentLocationsQueryBuilder) Count(ctx context.Context) (int, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get total")
	}
	return c, nil
}

// Collect creates iterator and collects all elements to slice.
func (b *GetRecentLocationsQueryBuilder) Collect(ctx context.Context) ([]Elem, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	r := make([]Elem, 0, c)
	for iter.Next(ctx) {
		r = append(r, iter.Value())
	}

	return r, iter.Err()
}

// GetRepliesQueryBuilder is query builder of MessagesGetReplies.
type GetRepliesQueryBuilder struct {
	raw        *tg.Client
	req        tg.MessagesGetRepliesRequest
	batchSize  int
	addOffset  int
	offsetDate int
	offsetID   int
}

// GetReplies creates query builder of MessagesGetReplies.
func (q *QueryBuilder) GetReplies(paramPeer tg.InputPeerClass) *GetRepliesQueryBuilder {
	b := &GetRepliesQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req: tg.MessagesGetRepliesRequest{
			Peer: &tg.InputPeerEmpty{},
		},
	}

	b.req.Peer = paramPeer
	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *GetRepliesQueryBuilder) BatchSize(batchSize int) *GetRepliesQueryBuilder {
	b.batchSize = batchSize
	return b
}

// OffsetDate sets offsetDate from which iterate start.
func (b *GetRepliesQueryBuilder) OffsetDate(offsetDate int) *GetRepliesQueryBuilder {
	b.offsetDate = offsetDate
	return b
}

// OffsetID sets offsetID from which iterate start.
func (b *GetRepliesQueryBuilder) OffsetID(offsetID int) *GetRepliesQueryBuilder {
	b.offsetID = offsetID
	return b
}

// MsgID sets MsgID field of GetReplies query.
func (b *GetRepliesQueryBuilder) MsgID(paramMsgID int) *GetRepliesQueryBuilder {
	b.req.MsgID = paramMsgID
	return b
}

// Peer sets Peer field of GetReplies query.
func (b *GetRepliesQueryBuilder) Peer(paramPeer tg.InputPeerClass) *GetRepliesQueryBuilder {
	b.req.Peer = paramPeer
	return b
}

// Query implements Query interface.
func (b *GetRepliesQueryBuilder) Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error) {
	r := &tg.MessagesGetRepliesRequest{
		Limit: req.Limit,
	}

	r.MsgID = b.req.MsgID
	r.Peer = b.req.Peer
	r.AddOffset = req.AddOffset
	r.OffsetDate = req.OffsetDate
	r.OffsetID = req.OffsetID
	return b.raw.MessagesGetReplies(ctx, r)
}

// Iter returns iterator using built query.
func (b *GetRepliesQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	iter = iter.OffsetDate(b.offsetDate)
	iter = iter.OffsetID(b.offsetID)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *GetRepliesQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}

// Count fetches remote state to get number of elements.
func (b *GetRepliesQueryBuilder) Count(ctx context.Context) (int, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get total")
	}
	return c, nil
}

// Collect creates iterator and collects all elements to slice.
func (b *GetRepliesQueryBuilder) Collect(ctx context.Context) ([]Elem, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	r := make([]Elem, 0, c)
	for iter.Next(ctx) {
		r = append(r, iter.Value())
	}

	return r, iter.Err()
}

// GetUnreadMentionsQueryBuilder is query builder of MessagesGetUnreadMentions.
type GetUnreadMentionsQueryBuilder struct {
	raw       *tg.Client
	req       tg.MessagesGetUnreadMentionsRequest
	batchSize int
	addOffset int
	offsetID  int
}

// GetUnreadMentions creates query builder of MessagesGetUnreadMentions.
func (q *QueryBuilder) GetUnreadMentions(paramPeer tg.InputPeerClass) *GetUnreadMentionsQueryBuilder {
	b := &GetUnreadMentionsQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req: tg.MessagesGetUnreadMentionsRequest{
			Peer: &tg.InputPeerEmpty{},
		},
	}

	b.req.Peer = paramPeer
	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *GetUnreadMentionsQueryBuilder) BatchSize(batchSize int) *GetUnreadMentionsQueryBuilder {
	b.batchSize = batchSize
	return b
}

// OffsetID sets offsetID from which iterate start.
func (b *GetUnreadMentionsQueryBuilder) OffsetID(offsetID int) *GetUnreadMentionsQueryBuilder {
	b.offsetID = offsetID
	return b
}

// Peer sets Peer field of GetUnreadMentions query.
func (b *GetUnreadMentionsQueryBuilder) Peer(paramPeer tg.InputPeerClass) *GetUnreadMentionsQueryBuilder {
	b.req.Peer = paramPeer
	return b
}

// Query implements Query interface.
func (b *GetUnreadMentionsQueryBuilder) Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error) {
	r := &tg.MessagesGetUnreadMentionsRequest{
		Limit: req.Limit,
	}

	r.Peer = b.req.Peer
	r.AddOffset = req.AddOffset
	r.OffsetID = req.OffsetID
	return b.raw.MessagesGetUnreadMentions(ctx, r)
}

// Iter returns iterator using built query.
func (b *GetUnreadMentionsQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	iter = iter.OffsetID(b.offsetID)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *GetUnreadMentionsQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}

// Count fetches remote state to get number of elements.
func (b *GetUnreadMentionsQueryBuilder) Count(ctx context.Context) (int, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get total")
	}
	return c, nil
}

// Collect creates iterator and collects all elements to slice.
func (b *GetUnreadMentionsQueryBuilder) Collect(ctx context.Context) ([]Elem, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	r := make([]Elem, 0, c)
	for iter.Next(ctx) {
		r = append(r, iter.Value())
	}

	return r, iter.Err()
}

// SearchQueryBuilder is query builder of MessagesSearch.
type SearchQueryBuilder struct {
	raw       *tg.Client
	req       tg.MessagesSearchRequest
	batchSize int
	addOffset int
	offsetID  int
}

// Search creates query builder of MessagesSearch.
func (q *QueryBuilder) Search(paramPeer tg.InputPeerClass) *SearchQueryBuilder {
	b := &SearchQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req: tg.MessagesSearchRequest{
			Filter: &tg.InputMessagesFilterEmpty{},
			FromID: &tg.InputPeerEmpty{},
			Peer:   &tg.InputPeerEmpty{},
		},
	}

	b.req.Peer = paramPeer
	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *SearchQueryBuilder) BatchSize(batchSize int) *SearchQueryBuilder {
	b.batchSize = batchSize
	return b
}

// OffsetID sets offsetID from which iterate start.
func (b *SearchQueryBuilder) OffsetID(offsetID int) *SearchQueryBuilder {
	b.offsetID = offsetID
	return b
}

// Filter sets Filter field of Search query.
func (b *SearchQueryBuilder) Filter(paramFilter tg.MessagesFilterClass) *SearchQueryBuilder {
	b.req.Filter = paramFilter
	return b
}

// FromID sets FromID field of Search query.
func (b *SearchQueryBuilder) FromID(paramFromID tg.InputPeerClass) *SearchQueryBuilder {
	b.req.FromID = paramFromID
	return b
}

// MaxDate sets MaxDate field of Search query.
func (b *SearchQueryBuilder) MaxDate(paramMaxDate int) *SearchQueryBuilder {
	b.req.MaxDate = paramMaxDate
	return b
}

// MinDate sets MinDate field of Search query.
func (b *SearchQueryBuilder) MinDate(paramMinDate int) *SearchQueryBuilder {
	b.req.MinDate = paramMinDate
	return b
}

// Peer sets Peer field of Search query.
func (b *SearchQueryBuilder) Peer(paramPeer tg.InputPeerClass) *SearchQueryBuilder {
	b.req.Peer = paramPeer
	return b
}

// Q sets Q field of Search query.
func (b *SearchQueryBuilder) Q(paramQ string) *SearchQueryBuilder {
	b.req.Q = paramQ
	return b
}

// TopMsgID sets TopMsgID field of Search query.
func (b *SearchQueryBuilder) TopMsgID(paramTopMsgID int) *SearchQueryBuilder {
	b.req.TopMsgID = paramTopMsgID
	return b
}

// ChatPhotos sets Filter field of Search query.
func (b *SearchQueryBuilder) ChatPhotos() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterChatPhotos{}
	return b
}

// Contacts sets Filter field of Search query.
func (b *SearchQueryBuilder) Contacts() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterContacts{}
	return b
}

// Document sets Filter field of Search query.
func (b *SearchQueryBuilder) Document() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterDocument{}
	return b
}

// Geo sets Filter field of Search query.
func (b *SearchQueryBuilder) Geo() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterGeo{}
	return b
}

// Gif sets Filter field of Search query.
func (b *SearchQueryBuilder) Gif() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterGif{}
	return b
}

// Music sets Filter field of Search query.
func (b *SearchQueryBuilder) Music() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterMusic{}
	return b
}

// MyMentions sets Filter field of Search query.
func (b *SearchQueryBuilder) MyMentions() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterMyMentions{}
	return b
}

// PhoneCalls sets Filter field of Search query.
func (b *SearchQueryBuilder) PhoneCalls(paramMissed bool) *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterPhoneCalls{
		Missed: paramMissed,
	}
	return b
}

// PhotoVideo sets Filter field of Search query.
func (b *SearchQueryBuilder) PhotoVideo() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterPhotoVideo{}
	return b
}

// Photos sets Filter field of Search query.
func (b *SearchQueryBuilder) Photos() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterPhotos{}
	return b
}

// Pinned sets Filter field of Search query.
func (b *SearchQueryBuilder) Pinned() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterPinned{}
	return b
}

// RoundVideo sets Filter field of Search query.
func (b *SearchQueryBuilder) RoundVideo() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterRoundVideo{}
	return b
}

// RoundVoice sets Filter field of Search query.
func (b *SearchQueryBuilder) RoundVoice() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterRoundVoice{}
	return b
}

// URL sets Filter field of Search query.
func (b *SearchQueryBuilder) URL() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterURL{}
	return b
}

// Video sets Filter field of Search query.
func (b *SearchQueryBuilder) Video() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterVideo{}
	return b
}

// Voice sets Filter field of Search query.
func (b *SearchQueryBuilder) Voice() *SearchQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterVoice{}
	return b
}

// Query implements Query interface.
func (b *SearchQueryBuilder) Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error) {
	r := &tg.MessagesSearchRequest{
		Limit: req.Limit,
	}

	r.Filter = b.req.Filter
	r.FromID = b.req.FromID
	r.MaxDate = b.req.MaxDate
	r.MinDate = b.req.MinDate
	r.Peer = b.req.Peer
	r.Q = b.req.Q
	r.TopMsgID = b.req.TopMsgID
	r.AddOffset = req.AddOffset
	r.OffsetID = req.OffsetID
	return b.raw.MessagesSearch(ctx, r)
}

// Iter returns iterator using built query.
func (b *SearchQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	iter = iter.OffsetID(b.offsetID)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *SearchQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}

// Count fetches remote state to get number of elements.
func (b *SearchQueryBuilder) Count(ctx context.Context) (int, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get total")
	}
	return c, nil
}

// Collect creates iterator and collects all elements to slice.
func (b *SearchQueryBuilder) Collect(ctx context.Context) ([]Elem, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	r := make([]Elem, 0, c)
	for iter.Next(ctx) {
		r = append(r, iter.Value())
	}

	return r, iter.Err()
}

// SearchGlobalQueryBuilder is query builder of MessagesSearchGlobal.
type SearchGlobalQueryBuilder struct {
	raw        *tg.Client
	req        tg.MessagesSearchGlobalRequest
	batchSize  int
	offsetID   int
	offsetPeer tg.InputPeerClass
	offsetRate int
}

// SearchGlobal creates query builder of MessagesSearchGlobal.
func (q *QueryBuilder) SearchGlobal() *SearchGlobalQueryBuilder {
	b := &SearchGlobalQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req: tg.MessagesSearchGlobalRequest{
			Filter: &tg.InputMessagesFilterEmpty{},
		},
	}

	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *SearchGlobalQueryBuilder) BatchSize(batchSize int) *SearchGlobalQueryBuilder {
	b.batchSize = batchSize
	return b
}

// OffsetID sets offsetID from which iterate start.
func (b *SearchGlobalQueryBuilder) OffsetID(offsetID int) *SearchGlobalQueryBuilder {
	b.offsetID = offsetID
	return b
}

// Filter sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Filter(paramFilter tg.MessagesFilterClass) *SearchGlobalQueryBuilder {
	b.req.Filter = paramFilter
	return b
}

// FolderID sets FolderID field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) FolderID(paramFolderID int) *SearchGlobalQueryBuilder {
	b.req.FolderID = paramFolderID
	return b
}

// MaxDate sets MaxDate field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) MaxDate(paramMaxDate int) *SearchGlobalQueryBuilder {
	b.req.MaxDate = paramMaxDate
	return b
}

// MinDate sets MinDate field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) MinDate(paramMinDate int) *SearchGlobalQueryBuilder {
	b.req.MinDate = paramMinDate
	return b
}

// Q sets Q field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Q(paramQ string) *SearchGlobalQueryBuilder {
	b.req.Q = paramQ
	return b
}

// ChatPhotos sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) ChatPhotos() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterChatPhotos{}
	return b
}

// Contacts sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Contacts() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterContacts{}
	return b
}

// Document sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Document() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterDocument{}
	return b
}

// Geo sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Geo() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterGeo{}
	return b
}

// Gif sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Gif() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterGif{}
	return b
}

// Music sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Music() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterMusic{}
	return b
}

// MyMentions sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) MyMentions() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterMyMentions{}
	return b
}

// PhoneCalls sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) PhoneCalls(paramMissed bool) *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterPhoneCalls{
		Missed: paramMissed,
	}
	return b
}

// PhotoVideo sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) PhotoVideo() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterPhotoVideo{}
	return b
}

// Photos sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Photos() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterPhotos{}
	return b
}

// Pinned sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Pinned() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterPinned{}
	return b
}

// RoundVideo sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) RoundVideo() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterRoundVideo{}
	return b
}

// RoundVoice sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) RoundVoice() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterRoundVoice{}
	return b
}

// URL sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) URL() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterURL{}
	return b
}

// Video sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Video() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterVideo{}
	return b
}

// Voice sets Filter field of SearchGlobal query.
func (b *SearchGlobalQueryBuilder) Voice() *SearchGlobalQueryBuilder {
	b.req.Filter = &tg.InputMessagesFilterVoice{}
	return b
}

// Query implements Query interface.
func (b *SearchGlobalQueryBuilder) Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error) {
	r := &tg.MessagesSearchGlobalRequest{
		Limit: req.Limit,
	}

	r.Filter = b.req.Filter
	r.FolderID = b.req.FolderID
	r.MaxDate = b.req.MaxDate
	r.MinDate = b.req.MinDate
	r.Q = b.req.Q
	r.OffsetID = req.OffsetID
	r.OffsetPeer = req.OffsetPeer
	r.OffsetRate = req.OffsetRate
	return b.raw.MessagesSearchGlobal(ctx, r)
}

// Iter returns iterator using built query.
func (b *SearchGlobalQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	iter = iter.OffsetID(b.offsetID)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *SearchGlobalQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}

// Count fetches remote state to get number of elements.
func (b *SearchGlobalQueryBuilder) Count(ctx context.Context) (int, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get total")
	}
	return c, nil
}

// Collect creates iterator and collects all elements to slice.
func (b *SearchGlobalQueryBuilder) Collect(ctx context.Context) ([]Elem, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	r := make([]Elem, 0, c)
	for iter.Next(ctx) {
		r = append(r, iter.Value())
	}

	return r, iter.Err()
}

// StatsGetMessagePublicForwardsQueryBuilder is query builder of StatsGetMessagePublicForwards.
type StatsGetMessagePublicForwardsQueryBuilder struct {
	raw        *tg.Client
	req        tg.StatsGetMessagePublicForwardsRequest
	batchSize  int
	offsetID   int
	offsetPeer tg.InputPeerClass
	offsetRate int
}

// StatsGetMessagePublicForwards creates query builder of StatsGetMessagePublicForwards.
func (q *QueryBuilder) StatsGetMessagePublicForwards(paramChannel tg.InputChannelClass) *StatsGetMessagePublicForwardsQueryBuilder {
	b := &StatsGetMessagePublicForwardsQueryBuilder{
		raw:       q.raw,
		batchSize: 1,
		req:       tg.StatsGetMessagePublicForwardsRequest{},
	}

	b.req.Channel = paramChannel
	return b
}

// BatchSize sets buffer of message loaded from one request.
// Be carefully, when set this limit, because Telegram does not return error if limit is too big,
// so results can be incorrect.
func (b *StatsGetMessagePublicForwardsQueryBuilder) BatchSize(batchSize int) *StatsGetMessagePublicForwardsQueryBuilder {
	b.batchSize = batchSize
	return b
}

// OffsetID sets offsetID from which iterate start.
func (b *StatsGetMessagePublicForwardsQueryBuilder) OffsetID(offsetID int) *StatsGetMessagePublicForwardsQueryBuilder {
	b.offsetID = offsetID
	return b
}

// Channel sets Channel field of StatsGetMessagePublicForwards query.
func (b *StatsGetMessagePublicForwardsQueryBuilder) Channel(paramChannel tg.InputChannelClass) *StatsGetMessagePublicForwardsQueryBuilder {
	b.req.Channel = paramChannel
	return b
}

// MsgID sets MsgID field of StatsGetMessagePublicForwards query.
func (b *StatsGetMessagePublicForwardsQueryBuilder) MsgID(paramMsgID int) *StatsGetMessagePublicForwardsQueryBuilder {
	b.req.MsgID = paramMsgID
	return b
}

// Query implements Query interface.
func (b *StatsGetMessagePublicForwardsQueryBuilder) Query(ctx context.Context, req Request) (tg.MessagesMessagesClass, error) {
	r := &tg.StatsGetMessagePublicForwardsRequest{
		Limit: req.Limit,
	}

	r.Channel = b.req.Channel
	r.MsgID = b.req.MsgID
	r.OffsetID = req.OffsetID
	r.OffsetPeer = req.OffsetPeer
	r.OffsetRate = req.OffsetRate
	return b.raw.StatsGetMessagePublicForwards(ctx, r)
}

// Iter returns iterator using built query.
func (b *StatsGetMessagePublicForwardsQueryBuilder) Iter() *Iterator {
	iter := NewIterator(b, b.batchSize)
	iter = iter.OffsetID(b.offsetID)
	return iter
}

// ForEach calls given callback on each iterator element.
func (b *StatsGetMessagePublicForwardsQueryBuilder) ForEach(ctx context.Context, cb func(context.Context, Elem) error) error {
	iter := b.Iter()
	for iter.Next(ctx) {
		if err := cb(ctx, iter.Value()); err != nil {
			return err
		}
	}
	return iter.Err()
}

// Count fetches remote state to get number of elements.
func (b *StatsGetMessagePublicForwardsQueryBuilder) Count(ctx context.Context) (int, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get total")
	}
	return c, nil
}

// Collect creates iterator and collects all elements to slice.
func (b *StatsGetMessagePublicForwardsQueryBuilder) Collect(ctx context.Context) ([]Elem, error) {
	iter := b.Iter()
	c, err := iter.Total(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	r := make([]Elem, 0, c)
	for iter.Next(ctx) {
		r = append(r, iter.Value())
	}

	return r, iter.Err()
}
