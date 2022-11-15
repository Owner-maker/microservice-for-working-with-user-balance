package models

import (
	"fmt"
)

func FillDbWithData() {
	user0 := User{Name: "Man Chelovekov"}
	DB.Create(&user0)

	fmt.Printf("User %s was added to the DB", user0.Name)
}
