package control

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	sTodaysSchedule		string = "Расписание на сегодня"
	sWeekSchedule 		string = "Расписание на неделю"
	sReminder			string = "Напоминание"
	sChangeScheduleDay 	string = "Изменить расписание"
	sReturnToMenu		string = "Вернуться назад"
	sMenuText			string = "Выберите опцию:"
	sDefaultText		string = "Я не понимаю эту команду."
	sReminderEnable		string = "Вкл/Выкл"
	sReminderChange		string = "Изменить время уведомления"
)

func HandleUpdate(update tgbotapi.Update) tgbotapi.MessageConfig {
	switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, sMenuText)
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton(sTodaysSchedule),
					tgbotapi.NewKeyboardButton(sWeekSchedule),
					tgbotapi.NewKeyboardButton(sReminder),
				),
			)
			return msg
		default:
			// Обработка нажатия на кнопки
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Text {
			case sTodaysSchedule:
				msg.Text = "*тут расписание на сегодня*"
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton(sChangeScheduleDay),
						tgbotapi.NewKeyboardButton(sReturnToMenu),
					),
				)
			case sWeekSchedule:
				msg.Text = "*тут расписание на неделю*"
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton(sChangeScheduleDay),
						tgbotapi.NewKeyboardButton(sReturnToMenu),
					),
				)
			case sReminder:
				msg.Text = "*тут описание напоминалки*"
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton(sReminderEnable),
						tgbotapi.NewKeyboardButton(sReminderChange),
						tgbotapi.NewKeyboardButton(sReturnToMenu),
					),
				)
			case sChangeScheduleDay:
				msg.Text = "*тут нужно ввести новое расписание на сегодня*"
			case sReturnToMenu:
				msg.Text = sMenuText
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton(sTodaysSchedule),
						tgbotapi.NewKeyboardButton(sWeekSchedule),
						tgbotapi.NewKeyboardButton(sReminder),
					),
				)
			default:
				msg.Text = sDefaultText
			}
		return msg
	}
}