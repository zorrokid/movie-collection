package main

import (
	"github.com/zorrokid/movieCollection/pkg/database"
)

func main() {
	_, err := database.InitDB()
	if err != nil {
		panic("failed connecting")
	}

}
