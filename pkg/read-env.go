package pkg

import (
	"log"
	file "path/filepath"

	e "mail.notification.com/config/env"
)

func GetEnv() *e.Env {
	path, err := file.Abs("./config/env")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	env := e.LoadEnv(path)
	return env
}
