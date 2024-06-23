package telegrambot

import (
	"backend/core/repos"
	"backend/core/services"
	"backend/pkg/logger"
	"backend/pkg/variable"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	EmployeeRequestService *services.EmployeeRequestService
	ClockService           *services.ClockService
	EmployeeRepo           *repos.EmployeeRepo
}

func NewBot() *Bot {
	return &Bot{
		EmployeeRequestService: services.NewEmployeeRequestService(),
		ClockService:           services.NewClockService(),
		EmployeeRepo:           repos.NewEmployeeRepo(),
	}
}

type BotInstance struct {
	bot *tgbotapi.BotAPI
}

var Instance *BotInstance

func (ctr *Bot) TelegramBot() {
	bottokken := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(bottokken)
	if err != nil {
		logger.Trace(err)
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
			employee, err := ctr.EmployeeRepo.FindTelegramId(&update.Message.From.ID)
			if err != nil {
				logger.Trace(err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is an error please contact the HR management.")
				Instance.bot.Send(msg)
				continue
			}
			if employee.ID == 0 {
				switch update.Message.Text {
				case "/start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome please register by chatting your name.")
					Instance.bot.Send(msg)
					continue
				default:
					ctr.AddToPending(update)
					continue
				}
			}

			// clock in or out with location
			if update.Message.Location != nil {
				var errMsg string
				var err error
				switch update.Message.ReplyToMessage.Text {
				case "Please send clock in location.":
					errMsg, err = ctr.ClockService.ClockIn(variable.Create[int](int(employee.ID)), update.Message.Location.Longitude, update.Message.Location.Latitude)
				case "Please send clock out location.":
					errMsg, err = ctr.ClockService.ClockOut(variable.Create[int](int(employee.ID)), update.Message.Location.Longitude, update.Message.Location.Latitude)
				default:
					continue
				}
				if err != nil {
					logger.Trace(err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is an error please contact the HR management.")
					Instance.bot.Send(msg)
					continue
				}
				var msg tgbotapi.MessageConfig
				if errMsg != "" {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, errMsg)
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Success.")
				}
				btn := tgbotapi.KeyboardButton{
					Text: "ClockIn",
				}
				btn2 := tgbotapi.KeyboardButton{
					Text: "ClockOut",
				}
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn, btn2})
				Instance.bot.Send(msg)
				continue
			}

			// getting clock out and in
			if update.Message.Text != "" {
				var msg string
				var errMsg string
				var err error
				switch update.Message.Text {
				case "ClockIn":
					errMsg, err = ctr.ClockService.CheckAvaiableClockIn(variable.Create[int](int(employee.ID)))
					msg = "Please send clock in location."
				case "ClockOut":
					errMsg, err = ctr.ClockService.CheckAvaiableClockOut(variable.Create[int](int(employee.ID)))
					msg = "Please send clock out location."
				case "Cancel":
					SendEmployeeAddedMessage(update.Message.From.ID)
				default:
					continue
				}
				if err != nil {
					logger.Trace(err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is an error please contact the HR management.")
					Instance.bot.Send(msg)
					continue
				}
				if errMsg != "" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, errMsg)
					Instance.bot.Send(msg)
					continue
				}
				msgConig := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
				btn := tgbotapi.KeyboardButton{
					RequestLocation: true,
					Text:            "Send Location",
				}
				btn2 := tgbotapi.KeyboardButton{
					Text: "Cancel",
				}
				msgConig.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn, btn2})
				Instance.bot.Send(msgConig)
				continue
			}
			// if update.Message.Text == "1" {
			// 	SendEmployeeAddedMessage(update.Message.From.ID)

			// }
		}
	}

}

func SendEmployeeAddedMessage(telegramID int64) {
	msg := tgbotapi.NewMessage(telegramID, "Hello you can call clock in and out from the button below.")
	btn := tgbotapi.KeyboardButton{
		Text: "ClockIn",
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
	ok, err = ctr.EmployeeRequestService.Pend(update.Message.Text, &update.Message.From.ID, update.Message.From.UserName)
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
