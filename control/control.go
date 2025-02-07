package control

import (
	"time"

	"github.com/adeptusmortem/workout-bot/database"
	"github.com/adeptusmortem/workout-bot/e"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Обработка сообщения
func HandleUpdate(update tgbotapi.Update) tgbotapi.MessageConfig {

	state, err := database.GetWaitingState(update.Message.Chat.ID)
	e.HandleError(err)

	switch state {
		case sStateNotFound:
			switch update.Message.Text {
				case "/start":
				database.CreateUser(update.Message.Chat.ID)
				database.ChangeWaitingState(update.Message.Chat.ID, sStateFree)
				return createMainMenu(update)
			default:
				return tgbotapi.NewMessage(update.Message.Chat.ID, "Напишите /start для начала пользовнаия ботом")
			}
		case sStateFree:
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case sStartCommand:
					return createMainMenu(update)
				default:
					return unknownMsg(update)
				}
			} else {
			switch update.Message.Text {
				case sTodaysSchedule:
					return createTodayMenu(update)
				case sWeekSchedule:
					return createWeakMenu(update)
				case sReminder:
					return createReminderMenu(update)
				case sChangeScheduleToday:
					return changeTodaySchedule(update)
				case sReturnToMenu:
					return createMainMenu(update)
				default:
					return unknownMsg(update)
				}
			}
		case sStateWaitingWeekdayShow:
			switch update.Message.Text {
				case sChangeScheduleDay:
					return changeDaySchedule(update)
				case sReturnToMenu:
					return createMainMenu(update)
				default:
					return showSchedule(update)
			}
		case sStateWaitingWeekdayChange:
			switch update.Message.Text {
				case sReturnToMenu:
					return createMainMenu(update)
				default:
					return makeNewSchedule(update)
			}
		case sStateWaitingSchedule:
			switch update.Message.Text {
				case sReturnToMenu:
					return createMainMenu(update)
				default:
					return handleNewSchedule(update)
			}
		default:
			database.ChangeWaitingState(update.Message.Chat.ID, sStateFree)
			return createMainMenu(update)

	}

}

// Создание основного меню
func createMainMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	database.ChangeWaitingState(update.Message.Chat.ID, sStateFree)
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
			tgbotapi.NewKeyboardButton(sChangeScheduleToday),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Создать меню изменения расписания на сегодня
func changeTodaySchedule(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = sAskNewWorkout
	database.ChangeWaitingState(update.Message.Chat.ID, sStateWaitingSchedule)
	database.ChangeWaitingParam(update.Message.Chat.ID, time.Now().Weekday().String())
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Создание меню расписаний на неделю
func createWeakMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	database.ChangeWaitingState(update.Message.Chat.ID, sStateWaitingWeekdayShow)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = sChooseWeekdayToShow
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		makeWeekdaysRow(),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sChangeScheduleDay),
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Показывает расписание на выбранный день
func showSchedule(update tgbotapi.Update) tgbotapi.MessageConfig {
	day, err := parseWeekday(update.Message.Text)
	e.HandleError(err)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if err != nil {
		return unknownMsg(update)
	}
	msg.Text, _ = database.GetWorkout(update.Message.Chat.ID, day)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Создать меню изменения расписания
func changeDaySchedule(update tgbotapi.Update) tgbotapi.MessageConfig {
	database.ChangeWaitingState(update.Message.Chat.ID, sStateWaitingWeekdayChange)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = sAskDay
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		makeWeekdaysRow(),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Создать меню ввода новой тренировки на выбранный день
func makeNewSchedule (update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	
	if _, exists := daysOfWeek[update.Message.Text]; !exists {
		msg.Text = sDaysScheduleUnchanged
		msg.ReplyMarkup = createMainMenu(update).ReplyMarkup
	}
	msg.Text = sAskNewWorkout
	database.ChangeWaitingState(update.Message.Chat.ID, sStateWaitingSchedule)
	database.ChangeWaitingParam(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Обработка нового расписания
func handleNewSchedule (update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ReplyMarkup = createMainMenu(update).ReplyMarkup

	sDay, err := database.GetWaitingParam(update.Message.Chat.ID)
	e.HandleError(err)

	database.ChangeWaitingState(update.Message.Chat.ID, sStateFree)

	day, err := parseWeekday(sDay)
	e.HandleError(err)

	if (err != nil) || (len(update.Message.Text) == 0) {
		msg.Text = sDaysScheduleUnchanged
		return msg
	}
	e.HandleError(database.ChangeDayWorkout(update.Message.Chat.ID, update.Message.Text, day))
	msg.Text = sDaysScheduleChanged
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
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(sReturnToMenu),
		),
	)
	return msg
}

// Обработка неизвестной команды
func unknownMsg(update tgbotapi.Update) tgbotapi.MessageConfig {
	database.ChangeWaitingState(update.Message.Chat.ID, sStateFree)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.Text == "/start" {
		return createMainMenu(update)
	}
	msg.Text = sDefaultText + " " + update.Message.Text
	return msg
}

// Создает ров кнопок дней недели
func makeWeekdaysRow() []tgbotapi.KeyboardButton {
	result := tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(Weekdays[0]),
			tgbotapi.NewKeyboardButton(Weekdays[1]),
			tgbotapi.NewKeyboardButton(Weekdays[2]),
			tgbotapi.NewKeyboardButton(Weekdays[3]),
			tgbotapi.NewKeyboardButton(Weekdays[4]),
			tgbotapi.NewKeyboardButton(Weekdays[5]),
			tgbotapi.NewKeyboardButton(Weekdays[6]),
	)
	return result
}