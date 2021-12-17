package handler

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func MakeNewMention(src *tb.Message) string {
	userNickname := src.Sender.Username
	userFirstName := src.Sender.FirstName
	userID := strconv.FormatInt(src.Sender.ID, 10)

	mention := "<a href=\"tg://user?id=" + userID + "\">@" + userNickname + "</a>"
	if userNickname == "" {
		mention = "<a href=\"tg://user?id=" + userID + "\">" + userFirstName + "</a>"
	}

	return mention
}
