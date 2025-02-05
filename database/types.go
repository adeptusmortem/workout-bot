package database

import "time"

type User struct {
	ID        int64 `gorm:"primaryKey"`
	Timezone  int8
	Schedules []Schedule `gorm:"foreignKey:UserID"`
	Reminders []Reminder `gorm:"foreignKey:UserID"`
}

type Schedule struct {
	ID          uint `gorm:"primaryKey"`
	UserID      int64
	DayOfWeek   time.Weekday
	WorkoutPlan string
}

type Reminder struct {
	ID          uint `gorm:"primaryKey"`
	UserID      int64
	Enabled     bool
	ScheduledAt time.Time
}