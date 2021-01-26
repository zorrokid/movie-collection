package main

import (
	"log"
	"os"

	"github.com/zorrokid/movieCollection/pkg/cvsloader"
	"github.com/zorrokid/movieCollection/pkg/database"
)

func main() {

	if len(os.Args) == 1 {
		panic("No input filename argument")
	}

	filename := os.Args[1]

	consts := database.GetDbConstants()

	log.Println(consts)

	err := cvsloader.LoadCollectionItems(filename, consts)

	if err != nil {
		log.Fatalln(err)
	}
}
