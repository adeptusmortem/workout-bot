package control

import (
	"errors"
	"strings"
	"time"

	"github.com/adeptusmortem/workout-bot/database"
	"github.com/adeptusmortem/workout-bot/e"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	sTodaysScheduleChanged	string = "Расписание на этот день изменено!"
	sDaysScheduleChanged	string = "Расписание на этот день изменено!"
)

// Карта дней недели
var daysOfWeek = map[string]time.Weekday{
	"sunday":		time.Sunday,
	"monday":		time.Monday,
	"tuesday":		time.Tuesday,
	"wednesday":	time.Wednesday,
	"thursday":		time.Thursday,
	"friday":		time.Friday,
	"saturday":		time.Saturday,
	"воскресенье":	time.Sunday,
	"понедельник":	time.Monday,
	"вторник":		time.Tuesday,
	"среда":		time.Wednesday,
	"четверг":		time.Thursday,
	"пятница":		time.Friday,
	"суббота":		time.Saturday,
	"вс":			time.Sunday,
	"пн":			time.Monday,
	"вт":			time.Tuesday,
	"ср":			time.Wednesday,
	"чт":			time.Thursday,
	"пт":			time.Friday,
	"сб":			time.Saturday,
}

// День недели из строки в time.Weekday
func parseWeekday(sDay string) (time.Weekday, error) {
	if day, exists := daysOfWeek[sDay]; exists {
		return day, nil
	}
    return time.Now().Weekday(), errors.New("unkonwn day of week")
}

func commandHandler (update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	// Парсинг сообщения на команду и текст
	parts := strings.Fields(update.Message.Text)
	if len(parts) == 0 {
		return unknownMsg(update)
	}

	command := strings.ToLower(strings.TrimPrefix(parts[0], "/"))	// Убрать / из команды и перевести в нижний регистр
	text := strings.Join(parts[1:], " ")			// Объединить оставшиеся части в текст
	sAnswerToCommand := "Не удалось обработать команду"

	switch command {
		case sChangeTodayCommand:
			sAnswerToCommand = handleCommandTodayChange(update.Message.Chat.ID, text)
		case sChangeDayCommand:
			parts = strings.Fields(update.Message.Text)
			sDay := strings.ToLower(parts[0])
			text = strings.Join(parts[1:], " ")
			sAnswerToCommand = handleCommandDayChange(update.Message.Chat.ID, text, sDay)
		default:
			return unknownMsg(update)
	}
	msg.Text = sAnswerToCommand

	return msg
}

// Обработчики команд
func handleCommandTodayChange(userID int64, text string) string {
	e.HandleError(database.ChangeTodayWorkout(userID, text))
	return sTodaysScheduleChanged
}

func handleCommandDayChange(userID int64, text string, sDay string) string {
	day, err := parseWeekday(sDay)
	e.HandleError(err)
	e.HandleError(database.ChangeDayWorkout(userID, text, day))
	return sDaysScheduleChanged
}