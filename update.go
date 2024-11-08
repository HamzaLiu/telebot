package telebot

import "strings"

// Update object represents an incoming update.
type Update struct {
	ID int `json:"update_id"`

	Message              *Message              `json:"message,omitempty"`
	EditedMessage        *Message              `json:"edited_message,omitempty"`
	ChannelPost          *Message              `json:"channel_post,omitempty"`
	EditedChannelPost    *Message              `json:"edited_channel_post,omitempty"`
	MessageReaction      *MessageReaction      `json:"message_reaction"`
	MessageReactionCount *MessageReactionCount `json:"message_reaction_count"`
	Callback             *Callback             `json:"callback_query,omitempty"`
	Query                *Query                `json:"inline_query,omitempty"`
	InlineResult         *InlineResult         `json:"chosen_inline_result,omitempty"`
	ShippingQuery        *ShippingQuery        `json:"shipping_query,omitempty"`
	PreCheckoutQuery     *PreCheckoutQuery     `json:"pre_checkout_query,omitempty"`
	Poll                 *Poll                 `json:"poll,omitempty"`
	PollAnswer           *PollAnswer           `json:"poll_answer,omitempty"`
	MyChatMember         *ChatMemberUpdate     `json:"my_chat_member,omitempty"`
	ChatMember           *ChatMemberUpdate     `json:"chat_member,omitempty"`
	ChatJoinRequest      *ChatJoinRequest      `json:"chat_join_request,omitempty"`
	Boost                *BoostUpdated         `json:"chat_boost"`
	BoostRemoved         *BoostRemoved         `json:"removed_chat_boost"`
}

// ProcessUpdate processes a single incoming update.
// A started bot calls this function automatically.
func (b *Bot) ProcessUpdate(u Update) {
	for _, handlers := range b.handlers {
		b.processUpdate(u, handlers)
	}
}

func (b *Bot) processUpdate(u Update, handlers []*Handler) {
	c := b.NewContext(u)

	if b.handle(OnAll, c, handlers) {
		return
	}

	if u.Message != nil {
		if b.handle(OnMessage, c, handlers) {
			return
		}

		m := u.Message

		if m.PinnedMessage != nil {
			b.handle(OnPinned, c, handlers)
			return
		}

		// Commands
		if m.Text != "" {
			// Filtering malicious messages
			if m.Text[0] == '\a' {
				return
			}

			match := cmdRx.FindAllStringSubmatch(m.Text, -1)
			if match != nil {
				// Syntax: "</command>@<bot> <payload>"
				command, botName := match[0][1], match[0][3]

				if botName != "" && !strings.EqualFold(b.Me.Username, botName) {
					return
				}

				m.Payload = match[0][5]
				if b.handle(command, c, handlers) {
					return
				}
			}

			// 1:1 satisfaction
			if b.handle(m.Text, c, handlers) {
				return
			}

			b.handle(OnText, c, handlers)
			return
		}

		if b.handleMedia(c, handlers) {
			return
		}

		if m.Contact != nil {
			b.handle(OnContact, c, handlers)
			return
		}
		if m.Location != nil {
			b.handle(OnLocation, c, handlers)
			return
		}
		if m.Venue != nil {
			b.handle(OnVenue, c, handlers)
			return
		}
		if m.Game != nil {
			b.handle(OnGame, c, handlers)
			return
		}
		if m.Dice != nil {
			b.handle(OnDice, c, handlers)
			return
		}
		if m.Invoice != nil {
			b.handle(OnInvoice, c, handlers)
			return
		}
		if m.Payment != nil {
			b.handle(OnPayment, c, handlers)
			return
		}

		if m.TopicCreated != nil {
			b.handle(OnTopicCreated, c, handlers)
			return
		}
		if m.TopicReopened != nil {
			b.handle(OnTopicReopened, c, handlers)
			return
		}
		if m.TopicClosed != nil {
			b.handle(OnTopicClosed, c, handlers)
			return
		}
		if m.TopicEdited != nil {
			b.handle(OnTopicEdited, c, handlers)
			return
		}
		if m.GeneralTopicHidden != nil {
			b.handle(OnGeneralTopicHidden, c, handlers)
			return
		}
		if m.GeneralTopicUnhidden != nil {
			b.handle(OnGeneralTopicUnhidden, c, handlers)
			return
		}
		if m.WriteAccessAllowed != nil {
			b.handle(OnWriteAccessAllowed, c, handlers)
			return
		}

		if m.TopicCreated != nil {
			b.handle(OnTopicCreated, c, handlers)
			return
		}
		if m.TopicReopened != nil {
			b.handle(OnTopicReopened, c, handlers)
			return
		}
		if m.TopicClosed != nil {
			b.handle(OnTopicClosed, c, handlers)
			return
		}
		if m.TopicEdited != nil {
			b.handle(OnTopicEdited, c, handlers)
			return
		}
		if m.GeneralTopicHidden != nil {
			b.handle(OnGeneralTopicHidden, c, handlers)
			return
		}
		if m.GeneralTopicUnhidden != nil {
			b.handle(OnGeneralTopicUnhidden, c, handlers)
			return
		}
		if m.WriteAccessAllowed != nil {
			b.handle(OnWriteAccessAllowed, c, handlers)
			return
		}

		wasAdded := (m.UserJoined != nil && m.UserJoined.ID == b.Me.ID) ||
			(m.UsersJoined != nil && isUserInList(b.Me, m.UsersJoined))
		if m.GroupCreated || m.SuperGroupCreated || wasAdded {
			b.handle(OnAddedToGroup, c, handlers)
			return
		}

		if m.UserJoined != nil {
			b.handle(OnUserJoined, c, handlers)
			return
		}
		if m.UsersJoined != nil {
			for _, user := range m.UsersJoined {
				m.UserJoined = &user
				b.handle(OnUserJoined, c, handlers)
			}
			return
		}
		if m.UserLeft != nil {
			b.handle(OnUserLeft, c, handlers)
			return
		}

		if m.UserShared != nil {
			b.handle(OnUserShared, c, handlers)
			return
		}
		if m.ChatShared != nil {
			b.handle(OnChatShared, c, handlers)
			return
		}

		if m.UserShared != nil {
			b.handle(OnUserShared, c, handlers)
			return
		}
		if m.ChatShared != nil {
			b.handle(OnChatShared, c, handlers)
			return
		}

		if m.NewGroupTitle != "" {
			b.handle(OnNewGroupTitle, c, handlers)
			return
		}
		if m.NewGroupPhoto != nil {
			b.handle(OnNewGroupPhoto, c, handlers)
			return
		}
		if m.GroupPhotoDeleted {
			b.handle(OnGroupPhotoDeleted, c, handlers)
			return
		}

		if m.GroupCreated {
			b.handle(OnGroupCreated, c, handlers)
			return
		}
		if m.SuperGroupCreated {
			b.handle(OnSuperGroupCreated, c, handlers)
			return
		}
		if m.ChannelCreated {
			b.handle(OnChannelCreated, c, handlers)
			return
		}

		if m.MigrateTo != 0 {
			m.MigrateFrom = m.Chat.ID
			b.handle(OnMigration, c, handlers)
			return
		}

		if m.VideoChatStarted != nil {
			b.handle(OnVideoChatStarted, c, handlers)
			return
		}
		if m.VideoChatEnded != nil {
			b.handle(OnVideoChatEnded, c, handlers)
			return
		}
		if m.VideoChatParticipants != nil {
			b.handle(OnVideoChatParticipants, c, handlers)
			return
		}
		if m.VideoChatScheduled != nil {
			b.handle(OnVideoChatScheduled, c, handlers)
			return
		}

		if m.WebAppData != nil {
			b.handle(OnWebApp, c, handlers)
			return
		}

		if m.ProximityAlert != nil {
			b.handle(OnProximityAlert, c, handlers)
			return
		}
		if m.AutoDeleteTimer != nil {
			b.handle(OnAutoDeleteTimer, c, handlers)
			return
		}
	}

	if u.EditedMessage != nil {
		b.handle(OnEdited, c, handlers)
		return
	}

	if u.ChannelPost != nil {
		m := u.ChannelPost

		if m.PinnedMessage != nil {
			b.handle(OnPinned, c, handlers)
			return
		}

		b.handle(OnChannelPost, c, handlers)
		return
	}

	if u.EditedChannelPost != nil {
		b.handle(OnEditedChannelPost, c, handlers)
		return
	}

	if u.Callback != nil {
		if data := u.Callback.Data; data != "" && data[0] == '\f' {
			match := cbackRx.FindAllStringSubmatch(data, -1)
			if match != nil {
				unique, payload := match[0][1], match[0][3]
				for _, handler := range handlers {
					if !handler.End.MatchString(unique) {
						continue
					}
					u.Callback.Unique = unique
					u.Callback.Data = payload
					b.runHandler(handler.HandlerFunc, c)
					return
				}
			}
		}

		b.handle(OnCallback, c, handlers)
		return
	}

	if u.Query != nil {
		b.handle(OnQuery, c, handlers)
		return
	}

	if u.InlineResult != nil {
		b.handle(OnInlineResult, c, handlers)
		return
	}

	if u.ShippingQuery != nil {
		b.handle(OnShipping, c, handlers)
		return
	}

	if u.PreCheckoutQuery != nil {
		b.handle(OnCheckout, c, handlers)
		return
	}

	if u.Poll != nil {
		b.handle(OnPoll, c, handlers)
		return
	}

	if u.PollAnswer != nil {
		b.handle(OnPollAnswer, c, handlers)
		return
	}

	if u.MyChatMember != nil {
		b.handle(OnMyChatMember, c, handlers)
		return
	}

	if u.ChatMember != nil {
		b.handle(OnChatMember, c, handlers)
		return
	}

	if u.ChatJoinRequest != nil {
		b.handle(OnChatJoinRequest, c, handlers)
		return
	}

	if u.Boost != nil {
		b.handle(OnBoost, c, handlers)
		return
	}

	if u.BoostRemoved != nil {
		b.handle(OnBoostRemoved, c, handlers)
		return
	}

	if u.Boost != nil {
		b.handle(OnBoost, c, handlers)
		return
	}

	if u.BoostRemoved != nil {
		b.handle(OnBoostRemoved, c, handlers)
		return
	}
}

func (b *Bot) handle(end string, c Context, handlers []*Handler) bool {
	for _, handler := range handlers {
		if handler.End.Match([]byte(end)) {
			b.runHandler(handler.HandlerFunc, c)
			return true
		}
	}
	return false
}

func (b *Bot) handleMedia(c Context, handlers []*Handler) bool {
	var (
		m     = c.Message()
		fired = true
	)

	switch {
	case m.Photo != nil:
		fired = b.handle(OnPhoto, c, handlers)
	case m.Voice != nil:
		fired = b.handle(OnVoice, c, handlers)
	case m.Audio != nil:
		fired = b.handle(OnAudio, c, handlers)
	case m.Animation != nil:
		fired = b.handle(OnAnimation, c, handlers)
	case m.Document != nil:
		fired = b.handle(OnDocument, c, handlers)
	case m.Sticker != nil:
		fired = b.handle(OnSticker, c, handlers)
	case m.Video != nil:
		fired = b.handle(OnVideo, c, handlers)
	case m.VideoNote != nil:
		fired = b.handle(OnVideoNote, c, handlers)
	default:
		return false
	}

	if !fired {
		return b.handle(OnMedia, c, handlers)
	}

	return true
}

func (b *Bot) runHandler(h HandlerFunc, c Context) {
	f := func() {
		if err := h(c); err != nil {
			b.OnError(err, c)
		}
	}
	if b.synchronous {
		f()
	} else {
		go f()
	}
}

func isUserInList(user *User, list []User) bool {
	for _, user2 := range list {
		if user.ID == user2.ID {
			return true
		}
	}
	return false
}
