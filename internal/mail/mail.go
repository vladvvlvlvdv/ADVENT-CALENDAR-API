package mail

import (
	"advent-calendar/internal/config"
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

	type ToSend struct {
		UnSubscribeLink string
		Year            int
		Title           string
		Text            string
		ClientLink      string
	}

	year := time.Now().Year()

	for email, toSend := range toSendEmails {
		link := fmt.Sprintf("%s/api/users/subscribe?email=%s", config.Config.MAIL_URI, email)
		clientLink := config.Config.CLIENT_URI

		if toSend.skipedDaysCount > 0 {
			title := "Не просмотренные дни - Кибербезопасный новый год"

			text := fmt.Sprintf(`Вы еще не просмотрели %d %s`,
				toSend.skipedDaysCount,
				utils.DeclOfNum(toSend.skipedDaysCount, [3]string{"рекомендацию", "рекомендации", "рекомендаций"}),
			)

			body, err := utils.LoadTemplate("message.email", ToSend{
				Title:           title,
				Text:            text,
				Year:            year,
				UnSubscribeLink: link,
				ClientLink:      clientLink,
			})
			if err != nil {
				log.Println("Error loading template:", err)
				continue
			}

			if err := utils.SendMail(email, title, body.String()); err != nil {
				log.Println("Error sending email to", email, ":", err)
				continue
			}
		}

		title := "Новая рекомендация - Кибербезопасный новый год"

		text := "Новый день - новая рекомендация! Не пропусти полезный совет"

		body, err := utils.LoadTemplate("message.email", ToSend{
			Title:           title,
			Text:            text,
			Year:            year,
			UnSubscribeLink: link,
			ClientLink:      clientLink,
		})
		if err != nil {
			log.Println("Error loading template:", err)
			continue
		}

		if err := utils.SendMail(email, title, body.String()); err != nil {
			log.Println("Error sending email to", email, ":", err)
			continue
		}
	}
}
