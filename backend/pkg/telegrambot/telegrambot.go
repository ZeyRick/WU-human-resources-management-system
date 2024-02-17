package telegrambot

import (
	"backend/core/services"
	"backend/core/types"
	"backend/pkg/logger"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	EmployeeRequestService *services.EmployeeRequestService
	ClockService           *services.ClockService
}

func NewBot() *Bot {
	return &Bot{
		EmployeeRequestService: services.NewEmployeeRequestService(),
		ClockService:           services.NewClockService(),
	}
}

type BotInstance struct {
	bot *tgbotapi.BotAPI
}

var Instance *BotInstance

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
			//log.Printf("[%s] \n %s", update.Message.Chat.ID, &update.Message.Text)
			if update.Message.Text == "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome if you have not registered please register by using /register [name].")
				Instance.bot.Send(msg)
			}
			if strings.Contains(update.Message.Text, "/register ") {
				ctr.AddToPending(update)
			}
			if update.Message.Location != nil {
				fmt.Println("1")
				err := ctr.ClockService.ClockFromTelegram(&update.Message.From.ID, types.ClockIn)
				if err != nil {
					logger.Trace(err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is an error please contact the HR management.")
					Instance.bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Clocked In.")
				Instance.bot.Send(msg)
			}
			if update.Message.Text == "ClockOut" {
				fmt.Println("2")
				err := ctr.ClockService.ClockFromTelegram(&update.Message.From.ID, types.ClockOut)
				if err != nil {
					logger.Trace(err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is an error please contact the HR management.")
					Instance.bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Clocked Out.")
				Instance.bot.Send(msg)
			}
			if update.Message.Text == "1" {
				SendEmployeeAddedMessage(update.Message.From.ID)
			}
		}
	}

}

func SendEmployeeAddedMessage(telegramID int64) {
	msg := tgbotapi.NewMessage(telegramID, "Hello you can call clock in and out from the button below.")
	btn := tgbotapi.KeyboardButton{
		RequestLocation: true,
		Text:            "ClockIn",
	}
	btn2 := tgbotapi.KeyboardButton{
		Text: "ClockOut",
	}
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn, btn2})
	Instance.bot.Send(msg)
}

func SendEmployeeRejectedMessage(telegramID int64) {
	msg := tgbotapi.NewMessage(telegramID, "Your register have been rejected.")
	Instance.bot.Send(msg)
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

func ButtonInChat(keyboardText string, callBackData string, keyboardText2 string, callBackData2 string) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(keyboardText, callBackData),
			tgbotapi.NewInlineKeyboardButtonData(keyboardText2, callBackData2),
		),
	)
	return keyboard
}

func (ctr *Bot) AddToPending(update tgbotapi.Update) {
	if update.Message.From.UserName == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your username is hidden or haven't added please try again")
		Instance.bot.Send(msg)
		return
	}
	ok, err := ctr.EmployeeRequestService.CheckPending(&update.Message.From.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is an error please contact the HR management")
		Instance.bot.Send(msg)
		return
	}
	if !ok {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your account has already registered")
		Instance.bot.Send(msg)
		return
	}
	ok, err = ctr.EmployeeRequestService.Pend(strings.Replace(update.Message.Text, "/register ", "", -1), &update.Message.From.ID, update.Message.From.UserName)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is an error please contact the HR management")
		Instance.bot.Send(msg)
		return
	}
	if !ok {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "This name has not added to the database please contact the HR management")
		Instance.bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your name has been added to the pending list")
	Instance.bot.Send(msg)
}
