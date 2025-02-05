package database

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const (
	// sNoScheduleFound = "Нет данных"
)

func GetTodayWorkout(userID int64) (string, error) {
	var schedule Schedule
	today := time.Now().Weekday()

	result := DB.Where("user_id = ? AND day_of_week = ?", userID, today).
		First(&schedule)

	if result.Error != nil {
		// return sNoScheduleFound, result.Error
		return makeSchedule(today), result.Error
	}

	return schedule.WorkoutPlan, nil
}

func ChangeTodayWorkout(userID int64, newWorkout string) error {
	return ChangeDayWorkout(userID, newWorkout, time.Now().Weekday())
}

func ChangeDayWorkout(userID int64, workoutPlan string, day time.Weekday) error {
    // Расписание на заданный день
    var schedule Schedule
    result := DB.Where("user_id = ? AND day_of_week = ?", userID, day).
        First(&schedule)

    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        // Расписание на выбраный день не найдено, создаем новое
        schedule = Schedule{
            UserID:     userID,
            DayOfWeek:  day,
            WorkoutPlan: workoutPlan,
        }
        result = DB.Create(&schedule)
    } else if result.Error != nil {
        // Другая ошибка
        return result.Error
    } else {
        // Расписание найдено, обновляем его
        schedule.WorkoutPlan = workoutPlan
        result = DB.Save(&schedule)
    }

    return result.Error
}