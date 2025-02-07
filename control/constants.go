package control

const (
	sTodaysSchedule      string = "Расписание на сегодня"
	sWeekSchedule        string = "Расписание на неделю"
	sChooseWeekdayToShow string = "Выберите на какой день показать расписание"
	sReminder            string = "Напоминание"
	sChangeScheduleDay   string = "Изменить расписание"
	sChangeScheduleToday string = "Изменить расписание на сегодня"
	sAskDay              string = "На какой день?"
	sAskNewWorkout       string = "Введите новое расписание"
	sReturnToMenu        string = "Вернуться назад"
	sMenuText            string = "Выберите опцию:"
	sDefaultText         string = "Я не понимаю эту команду."
	sReminderEnable      string = "Вкл/Выкл"
	sReminderChange      string = "Изменить время уведомления"

	sStartCommand string = "start"

	sTodaysScheduleChanged string = "Расписание на этот день изменено!"
	sDaysScheduleChanged   string = "Расписание на этот день изменено!"
	sDaysScheduleUnchanged string = "Не удалось изменить расписание!"

	sMonday    string = "Понедельник"
	sTuesday   string = "Вторник"
	sWednesday string = "Среда"
	sThursday  string = "Четверг"
	sFriday    string = "Пятница"
	sSaturday  string = "Суббота"
	sSunday    string = "Воскресенье"

	sStateNotFound             uint8 = 0
	sStateFree                 uint8 = 1
	sStateWaitingWeekdayChange uint8 = 2
	sStateWaitingSchedule      uint8 = 3
	sStateWaitingWeekdayShow   uint8 = 4
)

var (
	Weekdays = []string{sMonday, sTuesday, sWednesday, sThursday, sFriday, sSaturday, sSunday}
)