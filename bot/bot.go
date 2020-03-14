package bot

import(
	"github.com/era/pomobot/bot/controller"
	"github.com/era/pomobot/config"
	"github.com/era/pomobot/pomodoro"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var chat map[int]*tb.Chat

func Start() {
	pomobot := pomodoro.MemoryBot() // where we store all the pomodori
	chat = make(map[int]*tb.Chat) //maps userid -> chat

	telegramBot, err := tb.NewBot(tb.Settings{
		Token:  config.GetToken(),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	controller.Init(pomobot)
	router(telegramBot, routes())
	go sweeper(telegramBot, pomobot)
	telegramBot.Start()

}

func router (telegramBot *tb.Bot, routes map[string]func(message *tb.Message, telegramBot *tb.Bot)) {
	for k, v := range routes {
		// Creates local copy
		function := v
		command := k
		telegramBot.Handle(command, func(m *tb.Message) {
			chat[m.Sender.ID] = m.Chat
			log.Print("Received ", command)
			function(m, telegramBot)
		})
	}
}

func sweeper (bot *tb.Bot, storage pomodoro.PomoBot) {
	for true {
		time.Sleep(time.Second * 30)
		log.Print("Checking updates")
		updates, _ := storage.UpdateState()
		for k, v := range updates {
			log.Print("Going to update ", k)
			var message string
			if v == pomodoro.WORKING {
				message = "You should now work for 25 minutes"
			} else {
				message = "You should rest now for 5 minutes"
			}
			bot.Send(chat[k], message)
		}
	}
}
