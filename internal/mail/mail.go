package mail

import (
	"advent-calendar/internal/repository"
	"log"
)

func ScheduleSendEmailsToUsers() {
	users, err := repository.UserService.GetAll(repository.User{Role: "user"})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(users)

}
