package main

import (
	"fmt"
	"golang_app/app_1/app/controllers"
	"golang_app/app_1/app/models"
	// "fmt"
	// "golang_app/app_1/config"
	// "log"
)

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()

	fmt.Println("valid")
	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
