package internal

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

type ENV struct {
	ConfName  string
	ConfValue string
}

func ReadConf() ENV {
	var env ENV
	conf, err := os.ReadFile("./configs/.env")
	if err != nil {
		log.Fatal("the application can't read the config from .env file")
	}
	
	r, _ := regexp.Compile(`(?s)App.*?=.*?\n`)
	FoundedConfigs := r.FindAllString(string(conf), -1)
	fmt.Println(FoundedConfigs)

	return env
}
