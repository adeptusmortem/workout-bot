package database

import (
	"errors"

	"gorm.io/gorm"
)

// Создать запись пользователя если её нет
func CreateUser(userID int64) error {
    var user User

    result := DB.First(&user, userID)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        // Пользователь не найден, создаем нового
        user = User{
            ID:       userID,
            Timezone: +3,
        }
        result = DB.Create(&user)
        if result.Error != nil {
            return result.Error
        }
    } else if result.Error != nil {
        // Другая ошибка
        return result.Error
    }
	
    return nil
}

// Получить данные о пользователе
func GetUser(userID int64) (*User, error) {
	var user User
	err := DB.Preload("Schedules").Preload("Reminders").Preload("WaitingState").
		First(&user, userID).Error
	return &user, err
}