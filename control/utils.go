package control

import (
	"errors"
	"strings"
	"time"
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
	sDay = strings.ToLower(sDay)
	if day, exists := daysOfWeek[sDay]; exists {
		return day, nil
	}
    return time.Now().Weekday(), errors.New("unkonwn day of week")
}