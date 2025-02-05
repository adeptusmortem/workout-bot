package e

import "log"

func HandleError(err error) {
	if err != nil {
		log.Println("Ошибка:", err)
		return
	}
}