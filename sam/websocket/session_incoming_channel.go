package websocket

import (
	"context"

	"github.com/crusttech/crust/internal/auth"
	"github.com/crusttech/crust/sam/types"
	"github.com/crusttech/crust/sam/websocket/incoming"
	"github.com/crusttech/crust/sam/websocket/outgoing"
)

func (s *Session) channelJoin(ctx context.Context, p *incoming.ChannelJoin) error {
	// @todo: check access / can we join this channel (should be done by service layer)

	s.subs.Add(p.ChannelID)

	// Telling all subscribers of the channel we're joining that we are joining.
	var chJoin = &outgoing.ChannelJoin{
		ID:     p.ChannelID,
		UserID: uint64toa(auth.GetIdentityFromContext(ctx).Identity()),
	}

	// Send to all channel subscribers
	s.sendToAllSubscribers(chJoin, p.ChannelID)

	return nil
}

func (s *Session) channelPart(ctx context.Context, p *incoming.ChannelPart) error {
	// @todo: check access / can we part this channel? (should be done by service layer)

	// First, let's unsubscribe, so we don't hear echos
	s.subs.Delete(p.ChannelID)

	// This payload will tell everyone that we're parting from ALL channels
	var chPart = &outgoing.ChannelPart{
		ID:     p.ChannelID,
		UserID: uint64toa(auth.GetIdentityFromContext(ctx).Identity()),
	}

	s.sendToAllSubscribers(chPart, p.ChannelID)

	return nil
}

func (s *Session) channelList(ctx context.Context, p *incoming.Channels) error {
	channels, err := s.svc.ch.With(ctx).Find(&types.ChannelFilter{IncludeMembers: true})
	if err != nil {
		return err
	}

	// @todo count members for all channels

	return s.sendReply(payloadFromChannels(channels))
}

func (s *Session) channelCreate(ctx context.Context, p *incoming.ChannelCreate) (err error) {
	ch := &types.Channel{
		Type: types.ChannelTypePublic,
	}

	if p.Name != nil {
		ch.Name = *p.Name
	}

	if p.Topic != nil {
		ch.Topic = *p.Topic
	}

	if p.Type != nil {
		ch.Type = types.ChannelType(*p.Type)
	}

	ch, err = s.svc.ch.With(ctx).Create(ch)
	if err != nil {
		return err
	}

	// Explicitly subscribe to newly created channel
	s.subs.Add(uint64toa(ch.ID))

	// @todo this should go over all user's sessons and subscribe there as well

	// @todo load channel member count

	pl := payloadFromChannel(ch)

	if ch.Type == types.ChannelTypePublic {
		return s.sendToAll(pl)
	}

	// By default, just send reply to user
	return s.sendReply(pl)
}

func (s *Session) channelDelete(ctx context.Context, p *incoming.ChannelDelete) (err error) {
	err = s.svc.ch.With(ctx).Delete(parseUInt64(p.ChannelID))
	if err != nil {
		return err
	}

	return s.sendToAllSubscribers(&outgoing.ChannelDeleted{
		ID:     p.ChannelID,
		UserID: uint64toa(auth.GetIdentityFromContext(ctx).Identity()),
	}, p.ChannelID)
}

func (s *Session) channelUpdate(ctx context.Context, p *incoming.ChannelUpdate) error {
	ch, err := s.svc.ch.With(ctx).FindByID(parseUInt64(p.ID))
	if err != nil {
		return err
	}

	if p.Name != nil {
		ch.Name = *p.Name
	}

	if p.Topic != nil {
		ch.Topic = *p.Topic
	}

	if p.Type != nil {
		ch.Type = types.ChannelType(*p.Type)
	}

	ch, err = s.svc.ch.With(ctx).Update(ch)
	if err != nil {
		return err
	}

	// @todo load channel member count

	return s.sendToAllSubscribers(payloadFromChannel(ch), p.ID)
}
