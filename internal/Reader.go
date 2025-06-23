package internal

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type ENV struct {
	ConfName  string
	ConfValue string
}
func ReadConf() []ENV {
	var env []ENV
	conf, err := os.ReadFile("./configs/.env")
	if err != nil {
		log.Fatal("the application can't read the config from .env file")
	}

	configsByLine := strings.Split(string(conf), "\n")

	for _, value := range configsByLine {
		prettierLine := strings.TrimSpace(value)
		if strings.HasPrefix(prettierLine, "#") || prettierLine == "" {
			continue
		}
		r, _ := regexp.Compile(`^([A-Za-z_]+)=(.*)$`)
		FoundedConfig := r.FindStringSubmatch(prettierLine)

		// the Config Format
		envConfFormat := ENV{ ConfName: FoundedConfig[1], ConfValue: FoundedConfig[2] }
		env = append(env, envConfFormat)
	}
	return env
}
