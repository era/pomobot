package bot

import(
	"github.com/era/pomobot/bot/controller"
	tb "gopkg.in/tucnak/telebot.v2")

func routes() map[string]func(m *tb.Message, b *tb.Bot) {
	return map[string]func(m *tb.Message, b *tb.Bot){
		"/start": func(m *tb.Message, b *tb.Bot) {
			b.Send(m.Sender, controller.Instructions())
		},
		"/begin": func(m *tb.Message, b *tb.Bot) {
			b.Send(m.Sender, controller.CreateNew(m.Sender.ID))
		},
		"/end": func(m *tb.Message, b *tb.Bot) {
			b.Send(m.Sender, controller.End(m.Sender.ID))
		},
		"/pause": func(m *tb.Message, b *tb.Bot) {
			b.Send(m.Sender, controller.Pause(m.Sender.ID))
		},
		"/resume": func(m *tb.Message, b *tb.Bot) {
			b.Send(m.Sender, controller.Resume(m.Sender.ID))
		},
	}
}
