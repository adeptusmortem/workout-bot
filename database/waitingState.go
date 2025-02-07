package database

import (
	"errors"

	"gorm.io/gorm"
)

// Запросить статус
func GetWaitingState(userID int64) (uint8, error) {
	var waitingState WaitingState

	result := DB.Where("user_id = ?", userID).
		First(&waitingState)

	if result.Error != nil {
		return 0, result.Error // Добавить оработку ошибок?
	}

	return waitingState.State, nil
}

// Управление статусом ожидания ответа
func ChangeWaitingState(userID int64, state uint8) error {
    // Текущий статус
    var waitingState WaitingState
    result := DB.Where("user_id = ?", userID).
        First(&waitingState)

    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        // Записи нет, создаем новую
        waitingState = WaitingState{
            UserID:		userID,
            State: 		state,
            Param:      "",
        }
        result = DB.Create(&waitingState)
    } else if result.Error != nil {
        // Другая ошибка
        return result.Error
    } else {
        // Статус ожидания найден, обновляем его
        waitingState.State = state
        result = DB.Save(&waitingState)
    }

    return result.Error
}


// Запросить параметр
func GetWaitingParam(userID int64) (string, error) {
	var waitingState WaitingState

	result := DB.Where("user_id = ?", userID).
		First(&waitingState)

	if result.Error != nil {
		return "None", result.Error // Добавить оработку ошибок?
	}

	return waitingState.Param, nil
}

// Управление параметром ожидания ответа
func ChangeWaitingParam(userID int64, param string) error {
    // Текущий статус
    var waitingState WaitingState
    result := DB.Where("user_id = ?", userID).
        First(&waitingState)

    if result.Error != nil {
        // Ошибка
        return result.Error
    } else {
        // Статус ожидания найден, обновляем его
        waitingState.Param = param
        result = DB.Save(&waitingState)
    }

    return result.Error
}