package main

import (
	"fmt"
	"time"

	"test.com/first/project/models"
)

func main() {
	first := models.User{
		ID:        0,
		FirstName: "Edwin",
		LastName:  "Kuria",
	}

	duration := time.Duration(10) * time.Second

	fmt.Println(first)
	time.Sleep(duration)
}
