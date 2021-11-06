package handler

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func MakeNewMention(src *tb.Message) string {
	userNickname := src.Sender.Username
	userFirstName := src.Sender.FirstName
	userID := src.Sender.ID

	mention := "<a href=\"tg://user?id=" + strconv.Itoa(userID) + "\">@" + userNickname + "</a>"
	if userNickname == "" {
		mention = "<a href=\"tg://user?id=" + strconv.Itoa(userID) + "\">" + userFirstName + "</a>"
	}

	return mention
}
