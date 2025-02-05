package control

import (
	"github.com/adeptusmortem/workout-bot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	sTodaysSchedule		string = "Расписание на сегодня"
	sWeekSchedule 		string = "Расписание на неделю"
	sReminder			string = "Напоминание"
	sChangeScheduleDay 	string = "Изменить расписание"
	sClueToday			string = "Введите команду " + sChangeTodayCommand + " и новое расписание"
	sReturnToMenu		string = "Вернуться назад"
	sMenuText			string = "Выберите опцию:"
	sDefaultText		string = "Я не понимаю эту команду."
	sReminderEnable		string = "Вкл/Выкл"
	sReminderChange		string = "Изменить время уведомления"
	sStartCommand		string = "/start"
	sChangeTodayCommand	string = "/сегодня"
	sChangeDayCommand	string = "/изменить"
)

// Обработка сообщения
func HandleUpdate(update tgbotapi.Update) tgbotapi.MessageConfig {
	switch update.Message.Text {
		case sStartCommand:
			database.CreateUser(update.Message.Chat.ID)
			return createMainMenu(update)
		case sTodaysSchedule:
			return createTodayMenu(update)
		case sWeekSchedule:
			return createWeakMenu(update)
		case sReminder:
			return createReminderMenu(update)
		case sChangeScheduleDay:
			return changeDaySchedule(update)
		case sReturnToMenu:
			return createMainMenu(update)
		default:
			return commandHandler(update)
	}
}

// Создание основного меню
func createMainMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, sMenuText)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sTodaysSchedule),
			tgbotapi.NewKeyboardButton(sWeekSchedule),
			tgbotapi.NewKeyboardButton(sReminder),
		),
	)
	return msg
}

// Создание меню на сегодня
func createTodayMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text, _ = database.GetTodayWorkout(update.Message.Chat.ID)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sChangeScheduleDay),
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Создание меню на неделю
func createWeakMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "*тут расписание на неделю*"
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sChangeScheduleDay),
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Создать меню изменения расписания
func changeDaySchedule(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = sClueToday
	return msg
}

// Создание меню управления напоминалкой
func createReminderMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "*тут описание напоминалки*"
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sReminderEnable),
			tgbotapi.NewKeyboardButton(sReminderChange),
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Обработка неизвестной команды
func unknownMsg(update tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, sDefaultText + " " + update.Message.Text)
}
