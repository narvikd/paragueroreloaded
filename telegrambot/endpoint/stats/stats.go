package stats

import (
	"bytes"
	"fmt"
	"github.com/dustin/go-humanize"
	tb "gopkg.in/tucnak/telebot.v2"
	"paraguero_reloaded/telegrambot"
	"runtime"
	"strings"
	"text/tabwriter"
	"time"
)

var statsStartTime = time.Now()

func SendStats(bot *tb.Bot, src *tb.Message) {
	if strings.ToLower(src.Text) == "/stats" {
		chatID := tb.ChatID(src.Chat.ID)
		telegrambot.SendMessage(bot, chatID, getStats())
	}
}

func getStats() string {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	w := &tabwriter.Writer{}
	buf := &bytes.Buffer{}
	w.Init(buf, 0, 4, 0, ' ', 0)

	fmt.Fprintf(w, "Go: \t%s\n", runtime.Version())
	fmt.Fprintf(w, "Uptime: \t%s\n", getDurationString(time.Now().Sub(statsStartTime)))
	fmt.Fprintf(w, "Memory used: \t%s / %s (%s garbage collected)\n",
		humanize.Bytes(stats.Alloc), humanize.Bytes(stats.Sys), humanize.Bytes(stats.TotalAlloc))
	fmt.Fprintf(w, "Concurrent tasks: \t%s\n", humanize.Comma(int64(runtime.NumGoroutine())))
	_ = w.Flush()
	out := "<pre>" + buf.String() + "</pre>"
	return out
}

func getDurationString(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}
