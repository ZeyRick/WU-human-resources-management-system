package bot

import (
	"backend/core/services"
	"backend/pkg/logger"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	ConfirmationService *services.ConfirmationService
}

func NewBot() *Bot {
	return &Bot{
		ConfirmationService: services.NewConfirmationService(),
	}
}

type BotInstance struct {
	bot *tgbotapi.BotAPI
}

var Instance *BotInstance
var registering bool

func (ctr *Bot) TelegramBot() {
	bot, err := tgbotapi.NewBotAPI("6727294709:AAEi4reWROwsc5SkjY-DfurFR2pBB_I6eBM")
	if err != nil {
		logger.Trace(err)
		log.Panic(err)
		return
	}
	Instance = &BotInstance{bot: bot}
	Instance.bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := Instance.bot.GetUpdatesChan(u)
	go ctr.HandleUpdate(updates)
}

func (ctr *Bot) HandleUpdate(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] \n %s", update.Message.Chat.ID, &update.Message.Text)
			if update.Message.Text == "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome if you have not registered please you can click the button below to register.")
				msg.ReplyMarkup = ButtonInChat("Register", "Register")
				Instance.bot.Send(msg)
			}
			if registering {
				err := ctr.ConfirmationService.Pend(update.CallbackQuery.From.ID)
				if err != nil {
					logger.Trace(err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is an erro please contact the HR management.")
					Instance.bot.Send(msg)
					return
				}
				registering = false
			}
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello "+update.Message.From.FirstName+" "+update.Message.From.LastName+" you can clock in and out from the button below.")
			//msg.ReplyMarkup = ButtonOnTypingSection([]string{"ClockIn", "ClockOut"})
			//bot.Send(msg)
		}
		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data
			if data == "Register" {
				registering = true
			}
		}
	}
}

func ButtonOnTypingSection(buttonText []string) tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(buttonText[0]),
			tgbotapi.NewKeyboardButton(buttonText[1]),
		),
	)
	return keyboard
}

func ButtonInChat(keyboardText string, callBackData string) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyboardText, callBackData),
		),
	)
	return keyboard
}
