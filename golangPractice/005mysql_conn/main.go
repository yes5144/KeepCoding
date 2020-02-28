package main

// https://github.com/zxysilent/silent/blob/master/19%E5%B9%B4%E6%9A%91%E5%81%87%E8%A7%86%E9%A2%91/teach/control/article.go
import (
	"log"

	"github.com/yes5144/KeepCoding/golangPractice/005mysql_conn/models"
)

func main() {
	// select one
	u, _ := models.SelectUserOne(2)
	log.Println(u)

	// get all
	users, err := models.SelectUserAll()
	if err != nil {
		log.Printf("select all error, Failed code: %#v", err)
	}
	log.Println(users)
	log.Printf("Type: %T, value: %#v", users, users)

	// update
	mod := models.User{
		ID:   3,
		Name: "JavaScript",
		Age:  25,
	}
	err = models.UpdateUser(&mod)
	if err != nil {
		log.Printf("model insert , Failed code: %#v", err)
	}

	// insert
	mod2 := models.User{
		ID:   7,
		Name: "Shell",
		Age:  71,
	}
	err = models.InsertUser(&mod2)
	if err != nil {
		log.Printf("model insert error, Failed code: %#v", err)
	}

}
