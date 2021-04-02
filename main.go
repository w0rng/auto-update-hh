package main

import (
	"./utils"
	"fmt"
	"log"
	"time"
)

const CODE = "тут ключь, полученный после авторизации"

func main() {
	accessToken, refreshToken, err := utils.Auth(CODE)

	if err != nil {
		log.Fatal(err)
	}

	for {
		resumes, err := utils.GetAllResume(accessToken)
		if err != nil {
			log.Fatal(err)
		}

		for _, resume := range resumes {
			fmt.Println(resume)
			statusCode, err := utils.PublishResume(resume, accessToken)
			if err != nil {
				log.Fatal(err)
			}
			switch statusCode {
			case 204:
				log.Println("резюме обновлено")
			case 400:
				log.Println("Обновляю токен")
				accessToken, refreshToken, err = utils.Reauth(refreshToken)
				if err != nil {
					log.Fatal(err)
				}
			case 429:
				log.Println("обновление еще не доступно")
			default:
				log.Fatalf("что-то пошло не так...\n%d", statusCode)
			}
		}

		time.Sleep(time.Minute)
	}
}
