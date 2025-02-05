package main

import (
	"log"
	"os"

	"github.com/adeptusmortem/workout-bot/control"
	"github.com/adeptusmortem/workout-bot/database"
	"github.com/adeptusmortem/workout-bot/e"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Чтение токена из файла
	token, err := readTokenFromFile("token.txt")
	if err != nil {
		log.Fatalf("Ошибка при чтении токена из файла: %v", err)
	}

	// Создание бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	
    // Инициализация БД
    e.HandleError(database.Init())
    e.HandleError(database.AutoMigrate())

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			bot.Send(control.HandleUpdate(update))
		}
	}
}

// Чтение токена из файла
func readTokenFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}