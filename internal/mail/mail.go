package mail

import (
	"advent-calendar/internal/repository"
	"advent-calendar/pkg/utils"
	"fmt"
	"log"
	"slices"
	"time"
)

func ScheduleSendEmailsToUsers() {
	users, err := repository.UserService.GetAllSubscribes()
	if err != nil {
		log.Println(err)
		return
	}

	type mail struct {
		skipedDaysCount int
	}

	var toSendEmails = map[string]mail{}

	currentDay := time.Now().Day() - 1

	for _, user := range users {

		var skipedDaysCount int

		for i := 1; i < currentDay; i++ {
			if !slices.ContainsFunc(user.Days, func(day repository.Day) bool {
				return uint(i) < day.ID
			}) {
				skipedDaysCount++
			}
		}

		toSendEmails[user.Email] = mail{skipedDaysCount: skipedDaysCount}
	}

	for email, toSend := range toSendEmails {
		if toSend.skipedDaysCount > 0 {
			if err := utils.SendMail(email,
				"Кибербезопасный новый год",
				fmt.Sprintf(`Вы еще не просмотрели %d %s <br>`,
					toSend.skipedDaysCount,
					utils.DeclOfNum(toSend.skipedDaysCount, [3]string{"рекомендацию", "рекомендации", "рекомендаций"}),
				)); err != nil {
				log.Println("Error sending email to", email, ":", err)
				continue
			}
		}

		if err := utils.SendMail(email, "Кибербезопасный новый год", "Новый день - новая рекомендация! <br> Не пропусти полезный совет"); err != nil {
			log.Println("Error sending email to", email, ":", err)
			continue
		}
	}
}
