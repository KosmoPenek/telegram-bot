package model

import (
	"testing"

	mocks "github.com/KosmoPenek/telegram-bot/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartMessage_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("hello", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 123,
	})

	assert.NoError(t, err)
}

func Test_UnknownCommand_ShouldAnswerWithHelpMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("не знаю эту команду", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "some text",
		UserID: 123,
	})

	assert.NoError(t, err)
}
